import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'tide-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {

  constructor() { }
  users = []
  ngOnInit(): void {

  }
}
