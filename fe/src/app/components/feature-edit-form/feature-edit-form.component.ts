import { Component, Inject, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FeatureToggle, FTdataEdit } from 'src/app/shared/feature-toggle.model';

@Component({
  selector: 'app-feature-edit-form',
  templateUrl: './feature-edit-form.component.html',
  styleUrls: ["./feature-edit-form.component.sass"]
})
export class FeatureEditFormComponent implements OnInit {
  featureToggleForm: AbstractControl;

  constructor(
    @Inject(MAT_DIALOG_DATA) public input: FTdataEdit,
    public dialogRef: MatDialogRef<FeatureEditFormComponent>,
    private fb: FormBuilder,
  ) { }

  ngOnInit() {
    this.featureToggleForm = this.fb.group({
      displayName: [this.input.data.displayName],
      technicalName: [this.input.data.technicalName, Validators.required],
      expiresOn: [this.formatDateToString(this.input.data.expiresOn)],
      description: [this.input.data.description],
      isInverted: [this.input.data.isInverted, Validators.required],
      customerIds: [this.input.data.customerIds],
    });
  }

  onSubmit(): void {
    let updatedFT:FeatureToggle = {...this.input.data, ...this.featureToggleForm.value};    
    this.input.submitFn(updatedFT);
    this.dialogRef.close();
  }
  
  // Date format hack to conform to datepickers required format, yyyy-MM-dd.
  // ISO String 2021-02-10T20:07:33.657Z => 2021-02-10
  formatDateToString(date:Date):string {
    return date ? new Date(date).toISOString().substr(0,10) : ""
  }

}