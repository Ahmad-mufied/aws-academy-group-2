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
import { FilterDialogComponent } from '../filter-dialog/filter-dialog.component';
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
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  users: User[] = [];
  filteredUsers: User[] = [];
  displayedColumns: string[] = ['name', 'email', 'role', 'status', 'actions'];
  searchQuery: string = '';
  filters: { startDate?: Date; endDate?: Date; status?: string } = {}; // Tambah ini

  constructor(private userService: UserService, private dialog: MatDialog) {}

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
    const dialogRef = this.dialog.open(UserFormComponent, { data: user });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.updateUser(result);
      }
    });
  }

  viewDetails(user: User): void {
    this.dialog.open(UserDetailComponent, { data: user, width: '400px' });
  }

  assignProducts(user: User): void {
    const dialogRef = this.dialog.open(ProductAssignComponent, { data: user, width: '400px' });
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