import { Component, OnInit } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
@Component({
  selector: 'tide-log',
  templateUrl: './log.component.html',
  styleUrls: ['./log.component.scss']
})
export class LogComponent implements OnInit {

  constructor(public nd: NotebooksService) {
    this.users = this.nd.appLogs
  }

  ngOnInit(): void {
  }
  users=[
    {},
    {}
  ]
}
