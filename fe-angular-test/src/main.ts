import { bootstrapApplication } from '@angular/platform-browser';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideRouter, RouterOutlet } from '@angular/router';
import { routes } from './app/routes';
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet, 
    CommonModule   
  ],
  template: `<router-outlet></router-outlet>`
})
export class App {}

bootstrapApplication(App, {
  providers: [
    provideRouter(routes), 
    provideAnimations()
  ]
}).catch(err => console.error(err));
