import { Routes } from '@angular/router';
import { UserListComponent } from './pages/homepage/components/user-list/user-list.component';

export const routes: Routes = [
  { path: '', redirectTo: 'users', pathMatch: 'full' }, // Default ke /users
  { path: 'users', component: UserListComponent },
];
