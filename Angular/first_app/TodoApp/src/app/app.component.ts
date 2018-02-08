import { Component } from '@angular/core';

// Our component class
class AppComponent {
  title:string = 'hellp app';
}

// The title for the template and the acutal template itself; use templateUrl for external HTML file
@Component({
  selector: 'app',
  template: `<h1>{{ title }}</h1>`
})

export class AppComponent {
  title:string = 'hello app';
}
