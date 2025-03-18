import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { UserService } from '../../services/user.service';
import { User } from '../../models/user.model';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { UserFormComponent } from '../user-form/user-form.component';
import { UserDetailComponent } from '../user-detail/user-detail.component';
import { ProductAssignComponent } from '../product-assign/product-assign.component';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';
import { UserFilterComponent } from '../user-filter/user-filter.component';

@Component({
  selector: 'app-user-list',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatButtonModule,
    MatIconModule,
    MatInputModule,
    FormsModule,
    UserFilterComponent
  ],
  template: `
    <div class="container">
    <div class="header">
      <h1>User Management</h1>
      <div class="actions">
        <mat-form-field>
          <input matInput placeholder="Search users..." [(ngModel)]="searchQuery" (ngModelChange)="applyFilters()">
        </mat-form-field>
        <button mat-icon-button (click)="openFilterDialog()">
          <mat-icon>filter_list</mat-icon>
        </button>
        <button mat-raised-button color="primary" (click)="openUserForm()">Add User</button>
      </div>
    </div>

    <table mat-table [dataSource]="filteredUsers" class="mat-elevation-z8">
        <ng-container matColumnDef="name">
          <th mat-header-cell *matHeaderCellDef>Name</th>
          <td mat-cell *matCellDef="let user">{{user.name}}</td>
        </ng-container>

        <ng-container matColumnDef="email">
          <th mat-header-cell *matHeaderCellDef>Email</th>
          <td mat-cell *matCellDef="let user">{{user.email}}</td>
        </ng-container>

        <ng-container matColumnDef="role">
          <th mat-header-cell *matHeaderCellDef>Role</th>
          <td mat-cell *matCellDef="let user">{{user.role}}</td>
        </ng-container>

        <ng-container matColumnDef="status">
          <th mat-header-cell *matHeaderCellDef>Status</th>
          <td mat-cell *matCellDef="let user">{{user.status}}</td>
        </ng-container>

        <ng-container matColumnDef="actions">
          <th mat-header-cell *matHeaderCellDef>Actions</th>
          <td mat-cell *matCellDef="let user">
            <button mat-icon-button (click)="viewDetails(user)">
              <mat-icon>visibility</mat-icon>
            </button>
            <button mat-icon-button (click)="editUser(user)">
              <mat-icon>edit</mat-icon>
            </button>
            <button mat-icon-button (click)="assignProducts(user)">
              <mat-icon>assignment</mat-icon>
            </button>
            <button mat-icon-button color="warn" (click)="deleteUser(user)">
              <mat-icon>delete</mat-icon>
            </button>
          </td>
        </ng-container>

        <tr mat-header-row *matHeaderRowDef="displayedColumns"></tr>
        <tr mat-row *matRowDef="let row; columns: displayedColumns;"></tr>
      </table>
    </div>
    <ng-template #loading>
      <p>Loading users...</p>
    </ng-template>
  `,
  styles: [`
    .container {
      padding: 20px;
    }
    .header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
    }
    .actions {
      display: flex;
      gap: 16px;
      align-items: center;
    }
    table {
      width: 100%;
    }
  `]
})
export class UserListComponent implements OnInit {
  users: User[] = [];
  filteredUsers: User[] = [];
  displayedColumns: string[] = ['name', 'email', 'role', 'status', 'actions'];
  searchQuery: string = '';
  filters: { startDate?: Date; endDate?: Date; status?: string } = {}; // Tambah ini

  constructor(
    private userService: UserService,
    private dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.userService.getUsers().subscribe(users => {
      this.users = users;
      this.filteredUsers = users;
    });
  }

  openFilterDialog(): void {
    const dialogRef = this.dialog.open(UserFilterComponent, {
      data: this.filters,
      width: '400px'
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.filters = result;
        this.applyFilters();
      }
    });
  }

  applyFilters(): void {
    this.filteredUsers = this.users.filter(user => {
      const matchesSearch = this.searchQuery
        ? (user.name.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
           user.email.toLowerCase().includes(this.searchQuery.toLowerCase()))
        : true;
      const matchesStartDate = this.filters.startDate
        ? new Date(user.dob) >= new Date(this.filters.startDate)
        : true;
      const matchesEndDate = this.filters.endDate
        ? new Date(user.dob) <= new Date(this.filters.endDate)
        : true;
      const matchesStatus = this.filters.status
        ? user.status === this.filters.status
        : true;
      return matchesSearch && matchesStartDate && matchesEndDate && matchesStatus;
    });
  }

  // onSearch(event: Event): void {
  //   const query = (event.target as HTMLInputElement).value.toLowerCase();
  //   this.filteredUsers = this.users.filter(user =>
  //     user.name.toLowerCase().includes(query) ||
  //     user.email.toLowerCase().includes(query)
  //   );
  // }

  openUserForm(): void {
    const dialogRef = this.dialog.open(UserFormComponent);
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.addUser(result);
      }
    });
  }

  editUser(user: User): void {
    const dialogRef = this.dialog.open(UserFormComponent, {
      data: user
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.updateUser(result);
      }
    });
  }

  viewDetails(user: User): void {
    this.dialog.open(UserDetailComponent, {
      data: user,
      width: '400px'
    });
  }

  assignProducts(user: User): void {
    const dialogRef = this.dialog.open(ProductAssignComponent, {
      data: user,
      width: '400px'
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.updateUser({ ...user, products: result });
      }
    });
  }

  deleteUser(user: User): void {
    const dialogRef = this.dialog.open(ConfirmDialogComponent, {
      data: { message: `Are you sure you want to delete ${user.name} (ID: ${user.id})?` }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.deleteUser(user.id!);
      }
    });
  }
}