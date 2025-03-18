import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { User } from '../models/user.model';

@Injectable({ providedIn: 'root' })
export class UserService {
  private users: User[] = [];
  private usersSubject = new BehaviorSubject<User[]>([]);

  constructor() {
    this.loadUsersFromStorage();
    if (this.users.length === 0) {
      this.users = [{
        id: "1",
        name: 'Joe',
        email: 'joe@example.com',
        dob: new Date('1996-02-12'),
        role: 'Admin',
        registeredDate: new Date('2025-02-12'),
        status: 'Active',
        products: []
      }];
      this.saveUsersToStorage();
    }
    this.updateUsers();
  }

  private loadUsersFromStorage(): void {
    const stored = localStorage.getItem('users');
    this.users = stored ? JSON.parse(stored, (key, value) => {
      if (key === 'dob' || key === 'registeredDate') return new Date(value);
      return value;
    }) : [];
  }

  private saveUsersToStorage(): void {
    localStorage.setItem('users', JSON.stringify(this.users));
  }

  private updateUsers(): void {
    this.usersSubject.next([...this.users]);
    this.saveUsersToStorage();
  }

  getUsers(): Observable<User[]> {
    return this.usersSubject.asObservable();
  }

  addUser(user: User): void {
    user.id = (this.users.length + 1).toString();
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

  deleteUser(id: string): void {
    this.users = this.users.filter(user => user.id !== id);
    this.updateUsers();
  }

  // searchUsers(query: string): void {
  //   const filteredUsers = this.users.filter(user =>
  //     user.name.toLowerCase().includes(query.toLowerCase()) ||
  //     user.email.toLowerCase().includes(query.toLowerCase())
  //   );
  //   this.usersSubject.next(filteredUsers);
  // }
}