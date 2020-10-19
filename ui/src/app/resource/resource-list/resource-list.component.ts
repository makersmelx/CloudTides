import { Component, OnInit } from '@angular/core';
import { Observable, of } from 'rxjs';

import { Item, ResourceService } from '../resource.service';

@Component({
  selector: 'tide-resource-list',
  templateUrl: './resource-list.component.html',
  styleUrls: ['./resource-list.component.scss'],
})
export class ResourceListComponent implements OnInit {

  constructor(
    private resourceService: ResourceService,
  ) {

  }

  list$: Observable<Item[]> = of([]);
  opened = false;

  add (resource: Item) {
    this.resourceService.addItem(resource).subscribe(item => {
      this.refreshList();
    });
  }

  cancel() {
    this.opened = false;
  }

  ngOnInit() {
    this.refreshList();
  }

  refreshList() {
    this.list$ = this.resourceService.getList();
  }

}
