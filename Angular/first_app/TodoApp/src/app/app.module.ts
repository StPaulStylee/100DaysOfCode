import { AppComponent } from './app.component';

// This is where we make our components known
@NgModule({
  // this is where we include the components we want to use in the module
  declaration: [AppComponent],
  // this is the entry point component for our module, usually there is only one
  bootstrap: [AppComponent]
})
export class AppModule {};
