import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserListComponent } from './components/user-list/user-list.component';
// import { UserAddComponent } from './components/user-add/user-add.component';

export const routes: Routes = [
  { path: 'users', component: UserListComponent },
  // { path: 'users/add', component: UserAddComponent }, 
  { path: '', redirectTo: '/users', pathMatch: 'full' } // Redirect default
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
