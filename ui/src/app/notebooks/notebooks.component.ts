import { Component, OnInit, HostListener } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router'
import { NotebooksService } from './notebooks.service'

@Component({
  selector: 'tide-notebooks',
  templateUrl: './notebooks.component.html',
  styleUrls: ['./notebooks.component.scss']
})
export class NotebooksComponent implements OnInit {

  constructor(public router: Router, private nb: NotebooksService) {
    this.noteBook = this.nb
    this.instanceForm = this.nb.instanceForm
  }
  ngOnInit(): void {
  }
  noteBook: NotebooksService
  instanceForm: any
  cancel () {    
    this.noteBook.createInstanceFlag = false
  }
  create () {
    const str = 'HOME.NOTEBOOKS.Create'
    if (this.nb.createInstanceTitle === str) {
      this.nb.createNewApp().subscribe(data => {})
    } else {
      this.nb.modifyApp().subscribe(data => {})
    }
    this.noteBook.createInstanceFlag = false
  }
}
