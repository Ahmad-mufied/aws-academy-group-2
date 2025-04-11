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
import { ProductAssignComponent } from '../product-assign/product-assign.component';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';
import { UserFilterComponent } from '../user-filter/user-filter.component';
import { RouterModule, Router } from '@angular/router';
import { UserDetailComponent } from '../user-detail/user-detail.component';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatChipsModule } from '@angular/material/chips';

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
    RouterModule,
    MatProgressSpinnerModule,
    MatChipsModule
  ],
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  users: User[] = [];
  filteredUsers: User[] = [];
  displayedColumns: string[] = ['name', 'email', 'role', 'status', 'actions'];
  searchQuery: string = '';
  filters: { startDate?: Date; endDate?: Date; status?: string } = {};
  isLoading: boolean = true; // Tambah untuk loading state

  constructor(
    private userService: UserService,
    private dialog: MatDialog,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.loadUsers();
  }

  loadUsers(): void {
    this.isLoading = true;
    this.userService.getUsers().subscribe(users => {
      this.users = users;
      this.filteredUsers = users;
      this.applyFilters();
      this.isLoading = false;
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

  resetFilters(): void {
    this.searchQuery = '';
    this.filters = {};
    this.applyFilters();
  }

  openUserForm(): void {
    this.router.navigate(['/users/add']); // Sesuai rute baru
  }

  editUser(user: User): void {
    this.router.navigate(['/users/edit', user.id]); // Sesuai rute baru
  }

  viewDetails(user: User): void {
    this.dialog.open(UserDetailComponent, { data: user, width: '500px' });
  }

  assignProducts(user: User): void {
    const dialogRef = this.dialog.open(ProductAssignComponent, { data: user, width: '400px' });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.updateUser({ ...user, products: result });
        this.loadUsers(); // Refresh data setelah update
      }
    });
  }

  deleteUser(user: User): void {
    const dialogRef = this.dialog.open(ConfirmDialogComponent, {
      data: { message: `Are you sure you want to delete ${user.name} (ID: ${user.id})?` },
      width: '300px'
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.deleteUser(user.id!);
        this.loadUsers(); // Refresh data setelah delete
      }
    });
  }
}