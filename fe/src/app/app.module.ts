import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './shared/material.module';
import { ToggleContainerComponent } from './components/toggle-container/toggle-container.component';
import { DisplayToggleFeatureComponent } from './components/display-toggle-feature/display-toggle-feature.component';
import { FeatureEditFormComponent } from './components/feature-edit-form/feature-edit-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    ToggleContainerComponent,
    DisplayToggleFeatureComponent,
    FeatureEditFormComponent 
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MaterialModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  entryComponents: [FeatureEditFormComponent],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
