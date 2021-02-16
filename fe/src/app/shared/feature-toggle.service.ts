import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs/internal/BehaviorSubject';
import { FeatureToggle } from './feature-toggle.model';

@Injectable({
  providedIn: 'root'
})
export class FeatureToggleService {

  private _featuresToggles: FeatureToggle[];
  private featuresToggles:BehaviorSubject<FeatureToggle[]>;
  private baseURL = 'http://localhost:5000/api/featuretoggle';

  constructor(private http: HttpClient) {
    this.featuresToggles = new BehaviorSubject<FeatureToggle[]>(this._featuresToggles);
  }

  private emitFeaturesToggles() {
    this.featuresToggles.next(this._featuresToggles);
  }

  public subscribeToFeaturesToggleChanges() {
    return this.featuresToggles.asObservable();
  }

  public getFeaturesToggles() {
    // REST call
    this.http.get<FeatureToggle[]>(this.baseURL).subscribe(
      result => {
        this._featuresToggles = result === null ? [] : result;
        this.emitFeaturesToggles();
      }
    );
  }

  public addFeaturesToggle(newFT: FeatureToggle) {
    // REST call
    this.http.post<string>(this.baseURL, newFT).subscribe(
      newId => {
        newFT.id = newId;
        this._featuresToggles.push(newFT)
        this.emitFeaturesToggles();
      }
    );
  }

  public editFeaturesToggle(newFT: FeatureToggle) {
    // REST call
    this.http.put(this.baseURL+"/"+newFT.id, newFT).subscribe(
      result => {
        this._featuresToggles = this._featuresToggles.map(ft => ft.id === newFT.id ? newFT : ft)
        this.emitFeaturesToggles();
      }
    );
  }

  public deleteFeatureToggle(deleteFT: FeatureToggle) {
    this.http.delete(this.baseURL+"/"+deleteFT.id).subscribe(
      result => {
        this._featuresToggles = this._featuresToggles.filter(ft => ft.id !== deleteFT.id);
        this.emitFeaturesToggles();
      }
    );
  }

}
