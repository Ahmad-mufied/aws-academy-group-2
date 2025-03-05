package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/domain"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// ErrProductNotFound is returned when a product cannot be found
	ErrProductNotFound = errors.New("product not found")

	// ErrInvalidProductID is returned when an invalid product ID is provided
	ErrInvalidProductID = errors.New("invalid product ID")
)

// MongoProductRepository implements ProductRepository using MongoDB
type MongoProductRepository struct {
	collection *mongo.Collection
	timeout    time.Duration
}

// MongoProductRepositoryOptions configures the repository
type MongoProductRepositoryOptions struct {
	Timeout    time.Duration
	Collection string
}

// DefaultMongoOptions returns default options
func DefaultMongoOptions() *MongoProductRepositoryOptions {
	return &MongoProductRepositoryOptions{
		Timeout:    10 * time.Second,
		Collection: "product",
	}
}

// NewMongoProductRepository creates a new MongoDB product repository
func NewMongoProductRepository(db *mongo.Database, opts *MongoProductRepositoryOptions) (*MongoProductRepository, error) {
	if db == nil {
		return nil, errors.New("database connection is required")
	}

	if opts == nil {
		opts = DefaultMongoOptions()
	}

	collection := db.Collection(opts.Collection)

	// Create indexes for better query performance
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "product_id", Value: 1}}, // Index on product_id (unique)
			Options: options.Index().SetUnique(true),       // Ensure unique product_id
		},
	}

	// Create the indexes
	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()

	_, err := collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		return nil, fmt.Errorf("failed to create indexes: %w", err)
	}

	return &MongoProductRepository{
		collection: collection,
		timeout:    opts.Timeout,
	}, nil
}

// withTimeout creates a child context with timeout
func (r *MongoProductRepository) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if _, hasDeadline := ctx.Deadline(); hasDeadline {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, r.timeout)
}

func (r *MongoProductRepository) executeWithTimeout(ctx context.Context, fn func(context.Context) error) error {
	ctx, cancel := r.withTimeout(ctx)
	defer cancel()
	return fn(ctx)
}

// FindByID retrieves a product by its ID
func (r *MongoProductRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidProductID
	}

	var product domain.Product
	err := r.executeWithTimeout(ctx, func(ctx context.Context) error {
		filter := bson.M{"product_id": id}
		err := r.collection.FindOne(ctx, filter).Decode(&product)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return ErrProductNotFound
			}
			return fmt.Errorf("error finding product: %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Exists checks if a product exists by ID
func (r *MongoProductRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	if id == uuid.Nil {
		return false, ErrInvalidProductID
	}

	var count int64
	err := r.executeWithTimeout(ctx, func(ctx context.Context) error {
		filter := bson.M{"product_id": id}
		opts := options.Count().SetLimit(1)

		var err error
		count, err = r.collection.CountDocuments(ctx, filter, opts)
		if err != nil {
			return fmt.Errorf("error checking product existence: %w", err)
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// FindAll retrieves all products with optional filtering
func (r *MongoProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.executeWithTimeout(ctx, func(ctx context.Context) error {
		// Add options for performance
		findOptions := options.Find().
			SetSort(bson.D{{Key: "name", Value: 1}}).
			SetProjection(bson.M{"_id": 0}) // Exclude internal MongoDB _id field

		filterBson := bson.M{} // Define an empty filter or customize as needed
		cursor, err := r.collection.Find(ctx, filterBson, findOptions)
		if err != nil {
			return fmt.Errorf("error finding products: %w", err)
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var product domain.Product
			if err := cursor.Decode(&product); err != nil {
				return fmt.Errorf("error decoding product: %w", err)
			}

			products = append(products, &product)
		}

		if err := cursor.Err(); err != nil {
			return fmt.Errorf("error iterating products: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}

// Update modifies an existing product
func (r *MongoProductRepository) Update(ctx context.Context, product *domain.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}

	if product.ProductID == uuid.Nil {
		return ErrInvalidProductID
	}

	// Preserve created timestamp
	original, err := r.FindByID(ctx, product.ProductID)
	if err != nil {
		return err
	}

	// Update timestamp
	product.UpdatedAt = time.Now().UTC()
	product.CreatedAt = original.CreatedAt

	filter := bson.M{"product_id": product.ProductID}
	update := bson.M{"$set": product}

	return r.executeWithTimeout(ctx, func(ctx context.Context) error {
		result, err := r.collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return fmt.Errorf("error updating product: %w", err)
		}
		if result.MatchedCount == 0 {
			return ErrProductNotFound
		}
		return nil
	})
}

// Create adds a new product
func (r *MongoProductRepository) Create(ctx context.Context, product *domain.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}

	if product.Name == "" {
		return errors.New("product name is required")
	}

	if product.ProductID == uuid.Nil {
		product.ProductID = uuid.New()
	}

	// UTC + 7
	now := time.Now().UTC().Add(7 * time.Hour)
	if product.CreatedAt.IsZero() {
		product.CreatedAt = now
	}
	if product.UpdatedAt.IsZero() {
		product.UpdatedAt = now
	}

	exists, err := r.Exists(ctx, product.ProductID)
	if err != nil {
		return fmt.Errorf("error checking product existence: %w", err)
	}
	if exists {
		return fmt.Errorf("product with ID %s already exists", product.ProductID)
	}

	return r.executeWithTimeout(ctx, func(ctx context.Context) error {
		// Create the product
		_, err := r.collection.InsertOne(ctx, product)
		if err != nil {
			return fmt.Errorf("error creating product: %w", err)
		}

		return nil
	})
}

// Delete removes a product by ID
func (r *MongoProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidProductID
	}

	filter := bson.M{"product_id": id}

	return r.executeWithTimeout(ctx, func(ctx context.Context) error {
		result, err := r.collection.DeleteOne(ctx, filter)
		if err != nil {
			return fmt.Errorf("error deleting product: %w", err)
		}
		if result.DeletedCount == 0 {
			return ErrProductNotFound
		}
		return nil
	})
}
