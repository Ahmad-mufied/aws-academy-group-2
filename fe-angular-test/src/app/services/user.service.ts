import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { User } from '../models/user.model';

@Injectable({ providedIn: 'root' })
export class UserService {
  private users: User[] = [];
  private usersSubject = new BehaviorSubject<User[]>(this.loadUsersFromStorage());
  private currentId = this.getLastUserId();

  constructor() {
    if (this.users.length === 0) {
      this.users = [
        {
          id: this.generateUserId(),
          name: 'Joe',
          email: 'joe@example.com',
          dob: new Date('1996-02-12'),
          role: 'Admin',
          registeredDate: new Date(),
          status: 'Active',
          products: []
        }
      ];
      this.updateUsers();
    }
  }

  private loadUsersFromStorage(): User[] {
    const stored = localStorage.getItem('users');
    return stored
      ? JSON.parse(stored, (key, value) =>
          key === 'dob' || key === 'registeredDate' ? new Date(value) : value
        )
      : [];
  }

  private saveUsersToStorage(): void {
    localStorage.setItem('users', JSON.stringify(this.users));
  }

  private updateUsers(): void {
    this.usersSubject.next([...this.users]);
    this.saveUsersToStorage();
  }

  private getLastUserId(): number {
    const storedUsers = this.loadUsersFromStorage();
    if (storedUsers.length === 0) return 1;
    return Math.max(...storedUsers.map(user => user.id), 0) + 1;
  }

  private generateUserId(): number {
    return this.currentId++;
  }

  getUsers(): Observable<User[]> {
    return this.usersSubject.asObservable();
  }

  getUserById(userId: number): User | undefined {
    return this.users.find(user => user.id === userId);
  }

  getUserByIdSync(userId: number): User | null {
    return this.users.find(user => user.id === userId) || null;
  }

  addUser(user: User): void {
    user.id = this.generateUserId();
    user.registeredDate = new Date();
    this.users.push(user);
    this.updateUsers();
  }

  addUserSync(user: User): void {
    if (!user.name || !user.email) {
      throw new Error('Name and Email are required.');
    }
    this.addUser(user);
  }

  updateUser(user: User): void {
    const index = this.users.findIndex(u => u.id === user.id);
    if (index !== -1) {
      this.users[index] = user;
      this.updateUsers();
    }
  }

  updateUserSync(user: User): void {
    if (!user.id) {
      throw new Error('User ID is required.');
    }
    this.updateUser(user);
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

  resetUsers(): void {
    this.users = [];
    this.updateUsers();
  }
}
