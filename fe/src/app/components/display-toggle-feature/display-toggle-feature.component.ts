import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { FeatureToggle, FTdataEdit } from 'src/app/shared/feature-toggle.model';
import { FeatureToggleService } from 'src/app/shared/feature-toggle.service';
import { FeatureEditFormComponent } from '../feature-edit-form/feature-edit-form.component';

@Component({
  selector: 'app-display-toggle-feature',
  templateUrl: './display-toggle-feature.component.html',
  styleUrls: ["./display-toggle-feature.component.sass"]
})
export class DisplayToggleFeatureComponent implements OnInit {
  constructor(public dialog: MatDialog,
              private ftService: FeatureToggleService
    ) { }
  @Input() inputData: FeatureToggle;
  
  onToggle() {
    this.inputData.isActive =! this.inputData.isActive;
    this.ftService.editFeaturesToggle(this.inputData);
  }

  openEditDialog() {
    var dataObject:FTdataEdit = {
      data: this.inputData,
      submitFn:(ft) => this.ftService.editFeaturesToggle(ft)
    };

    const dialogRef = this.dialog.open(FeatureEditFormComponent, {data:dataObject});

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

  archiveFeature() {
    this.ftService.deleteFeatureToggle(this.inputData);
  }

  ngOnInit() {
  }

}