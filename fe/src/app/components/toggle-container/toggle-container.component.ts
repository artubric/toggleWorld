import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Observable } from 'rxjs/internal/Observable';
import { FeatureToggle, FTdataEdit } from 'src/app/shared/feature-toggle.model';
import { FeatureToggleService } from 'src/app/shared/feature-toggle.service';
import { FeatureEditFormComponent } from '../feature-edit-form/feature-edit-form.component';

@Component({
  selector: 'app-toggle-container',
  templateUrl: './toggle-container.component.html',
  styleUrls: ["./toggle-container.component.sass"]
})
export class ToggleContainerComponent implements OnInit {
  public featureToggles$:Observable<FeatureToggle[]>;

  constructor(private featureToggleService: FeatureToggleService,
              public dialog: MatDialog) { }
  onCreateClick() {
    var dummyData:FeatureToggle = {
      id: "",
      isActive: false,
      technicalName: "new.feature.toggle",
      isInverted: false,
      customerIds: []
    }

    var dataObject:FTdataEdit = {
        data: dummyData, 
        submitFn:(ft:FeatureToggle) => this.featureToggleService.addFeaturesToggle(ft)
    }

    const dialogRef = this.dialog.open(FeatureEditFormComponent, {data: dataObject});

    dialogRef.afterClosed().subscribe(result => {
      //show snackbar?
      console.log(`Dialog result: ${result}`);
    });
  }

  ngOnInit() {
    this.featureToggles$ = this.featureToggleService.subscribeToFeaturesToggleChanges();
    this.featureToggleService.getFeaturesToggles();
  }

}
