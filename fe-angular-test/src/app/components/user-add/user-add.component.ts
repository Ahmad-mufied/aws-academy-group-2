import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../../services/user.service';
import { User } from '../../models/user.model';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-user-add',
  standalone: true,
  imports: [CommonModule, MatButtonModule, MatInputModule, FormsModule],
  templateUrl: './user-add.component.html',
  styleUrls: ['./user-add.component.css']
})
export class UserAddComponent {
  user: User = {
    name: '',
    email: '',
    role: 'Staff',
    status: 'Active',
    products: [],
    registeredDate: new Date(),
    dob: new Date(),
    id: '0'
  };

  constructor(public userService: UserService, public router: Router) {}

  onSave(): void {
    this.userService.addUser(this.user);
    this.router.navigate(['/users']);
  }
}