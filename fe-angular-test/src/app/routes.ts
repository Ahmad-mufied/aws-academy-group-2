import { Routes, RouterModule } from '@angular/router';
import { UserListComponent } from './components/user-list/user-list.component';
import { UserAddComponent } from './components/user-add/user-add.component';
import { UserFormComponent } from './components/user-form/user-form.component';

export const routes: Routes = [
  { path: 'users', component: UserListComponent },
  { path: 'add', component: UserAddComponent }, 
  { path: 'edit/:id', component: UserFormComponent },
  { path: '', redirectTo: '/users', pathMatch: 'full' } // Redirect default
];

export const AppRoutingModule = RouterModule.forRoot(routes);
