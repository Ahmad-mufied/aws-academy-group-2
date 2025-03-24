import { Routes, RouterModule } from '@angular/router';
import { UserListComponent } from './components/user-list/user-list.component';
import { UserFormComponent } from './components/user-form/user-form.component';

export const routes: Routes = [
  { path: 'users', component: UserListComponent },
  { path: 'users/add', component: UserFormComponent }, // Rute untuk Add
  { path: 'users/edit/:id', component: UserFormComponent }, // Rute untuk Edit
  { path: '', redirectTo: '/users', pathMatch: 'full' }
];

export const AppRoutingModule = RouterModule.forRoot(routes);