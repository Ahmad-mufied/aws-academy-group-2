// user-list.component.ts
import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatTableModule } from '@angular/material/table';
import { MatListModule } from '@angular/material/list';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatPaginator } from '@angular/material/paginator';
import { MatCardModule } from '@angular/material/card';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { UserDetailDialogComponent } from '../user-detail-dialog/user-detail-dialog.component';
import { AssignProductsDialogComponent } from '../assign-products-dialog/assign-products-dialog.component';

export interface User {
  id: number;
  name: string;
  email: string;
  dob: string;
  role: 'Admin' | 'Supervisor' | 'Manager' | 'Staff';
  registeredDate: string;
  status: 'Active' | 'Inactive';
}

@Component({
  selector: 'app-user-list',
  standalone: true,
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css'],
  imports: [
    MatTableModule,
    MatListModule,
    CommonModule,
    MatCheckboxModule,
    FormsModule,
    MatIconModule,
    MatDividerModule,
    MatButtonModule,
    MatPaginatorModule,
    MatCardModule,
    MatToolbarModule,
    MatTooltipModule,
    MatDialogModule,
  ],
})
export class UserListComponent implements OnInit, AfterViewInit {
  displayedColumns: string[] = ['name', 'email', 'dob', 'role', 'registeredDate', 'status', 'actions'];
  users: User[] = [];
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  constructor(public dialog: MatDialog) {}

  ngOnInit(): void {
    // Mock Data
    this.users = [
      { id: 1, name: 'John Doe', email: 'john.doe@example.com', dob: '1995-06-15', role: 'Admin', registeredDate: '2024-02-10', status: 'Active' },
      { id: 2, name: 'Jane Smith', email: 'jane.smith@example.com', dob: '1998-09-22', role: 'Staff', registeredDate: '2023-11-01', status: 'Inactive' },
      { id: 3, name: 'Alice Johnson', email: 'alice.johnson@example.com', dob: '1990-03-10', role: 'Manager', registeredDate: '2023-10-20', status: 'Active' },
      { id: 4, name: 'Bob Williams', email: 'bob.williams@example.com', dob: '1987-12-05', role: 'Supervisor', registeredDate: '2023-09-15', status: 'Inactive' },
      { id: 5, name: 'Charlie Brown', email: 'charlie.brown@example.com', dob: '1993-08-28', role: 'Staff', registeredDate: '2023-08-01', status: 'Active' },
      { id: 6, name: 'Diana Miller', email: 'diana.miller@example.com', dob: '1996-05-01', role: 'Admin', registeredDate: '2023-07-10', status: 'Inactive' },
      { id: 7, name: 'Ethan Davis', email: 'ethan.davis@example.com', dob: '1989-02-14', role: 'Manager', registeredDate: '2023-06-05', status: 'Active' },
      { id: 8, name: 'Fiona Wilson', email: 'fiona.wilson@example.com', dob: '1991-11-30', role: 'Supervisor', registeredDate: '2023-05-20', status: 'Inactive' },
      { id: 9, name: 'George Garcia', email: 'george.garcia@example.com', dob: '1994-07-08', role: 'Staff', registeredDate: '2023-04-15', status: 'Active' },
      { id: 10, name: 'Hannah Rodriguez', email: 'hannah.rodriguez@example.com', dob: '1986-04-22', role: 'Admin', registeredDate: '2023-03-01', status: 'Inactive' }
    ];
  }

  ngAfterViewInit() {
    this.paginator.pageSize = 5;
    this.paginator.pageSizeOptions = [5, 10, 25, 100];
  }

  viewUser(user: User) {
    const dialogRef = this.dialog.open(UserDetailDialogComponent, {
      width: '400px',
      data: user,
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
    });
  }

  editUser(user: User) {
    console.log('Editing user:', user);
  }

  assignRole(user: User) {
    const dialogRef = this.dialog.open(AssignProductsDialogComponent, {
      width: '500px',
      data: user,
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The assign products dialog was closed', result);
    });
  }

  deleteUser(user: User) {
    console.log('Deleting user:', user);
  }
}
