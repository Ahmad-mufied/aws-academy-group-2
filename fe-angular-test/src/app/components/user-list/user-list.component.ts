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

@Component({
  selector: 'app-user-list',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatButtonModule,
    MatIconModule,
    MatInputModule,
    FormsModule
  ],
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  users: User[] = [];
  allUsers: User[] = [];
  displayedColumns: string[] = ['name', 'email', 'role', 'status', 'actions'];

  constructor(private userService: UserService, private dialog: MatDialog) {}

  ngOnInit(): void {
    this.userService.getUsers().subscribe(users => {
      this.users = users;
      this.allUsers = users;
    });
  }

  onSearch(event: Event): void {
    const query = (event.target as HTMLInputElement).value.toLowerCase();
    this.users = this.allUsers.filter(user =>
      user.name.toLowerCase().includes(query) ||
      user.email.toLowerCase().includes(query)
    );
  }

  openFilterDialog(): void {
    const dialogRef = this.dialog.open(FilterDialogComponent, { width: '400px' });
    dialogRef.afterClosed().subscribe(filterData => {
      if (filterData) {
        this.applyFilter(filterData);
      }
    });
  }

  applyFilter(filterData: any): void {
    this.users = this.allUsers.filter(user => {
      const matchesStatus = filterData.status === 'All' || user.status === filterData.status;
      const matchesDOB = !filterData.dob || user.dob?.toISOString().slice(0, 10) === filterData.dob;
      return matchesStatus && matchesDOB;
    });
  }

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
    if (confirm(`Are you sure you want to delete ${user.name}?`)) {
      if (user.id !== undefined) {
        this.userService.deleteUser(user.id.toString());
      }
    }
  }
}