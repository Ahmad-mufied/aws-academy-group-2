import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatTableModule } from '@angular/material/table';
import { MatListModule } from '@angular/material/list';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon';


export interface User {
  name: string;
  email: string;
  birthDate: string;
  joinedAt: string;
  status: 'Active' | 'Inactive';
  products: { basic: boolean; inter: boolean; adv: boolean };
}

@Component({
  selector: 'app-user-list',
  standalone: true,
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss'],
  imports: [
    MatTableModule,
    MatListModule,
    CommonModule,
    MatCheckboxModule,
    FormsModule,
    MatIconModule
  ],
})
export class UserListComponent {
  displayedColumns: string[] = ['no', 'name', 'email', 'birthDate', 'joinedAt', 'status', 'product', 'actions'];
  users: User[] = [
    {
      name: 'John Doe',
      email: 'john@example.com',
      birthDate: '1995-06-15',
      joinedAt: '2024-02-10',
      status: 'Active',
      products: { basic: true, inter: true, adv: false }
    },
    {
      name: 'Jane Smith',
      email: 'jane@example.com',
      birthDate: '1998-09-22',
      joinedAt: '2023-11-01',
      status: 'Inactive',
      products: { basic: false, inter: true, adv: true }
    }
  ];

  viewUser(user: User) {
    console.log('Viewing user:', user);
  }

  assignProduct(user: User) {
    console.log('Assigning product to user:', user);
  }

  editUser(user: User) {
    console.log('Editing user:', user);
  }

  deleteUser(user: User) {
    console.log('Deleting user:', user);
  }
}
