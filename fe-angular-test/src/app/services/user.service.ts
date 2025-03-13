import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { User } from '../models/user.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private users: User[] = [];
  private usersSubject = new BehaviorSubject<User[]>([]);

  constructor() {
    // Initial dummy data
    this.users = [{
      id: 1,
      name: 'Joe',
      email: 'joe@example.com',
      dob: new Date('1996-02-12'),
      role: 'Admin',
      registeredDate: new Date('2025-02-12'),
      status: 'Active',
      products: []
    }];
    this.updateUsers();
  }

  private updateUsers(): void {
    this.usersSubject.next([...this.users]);
  }

  getUsers(): Observable<User[]> {
    return this.usersSubject.asObservable();
  }

  addUser(user: User): void {
    user.id = this.users.length + 1;
    this.users.push(user);
    this.updateUsers();
  }

  updateUser(user: User): void {
    const index = this.users.findIndex(u => u.id === user.id);
    if (index !== -1) {
      this.users[index] = user;
      this.updateUsers();
    }
  }

  deleteUser(id: number): void {
    this.users = this.users.filter(user => user.id !== id);
    this.updateUsers();
  }

  searchUsers(query: string): void {
    const filteredUsers = this.users.filter(user =>
      user.name.toLowerCase().includes(query.toLowerCase()) ||
      user.email.toLowerCase().includes(query.toLowerCase())
    );
    this.usersSubject.next(filteredUsers);
  }
}