import { Component, OnInit, HostListener } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router'
import { NotebooksService } from './notebooks.service'
import { FormBuilder, Validators } from '@angular/forms'

@Component({
  selector: 'tide-notebooks',
  templateUrl: './notebooks.component.html',
  styleUrls: ['./notebooks.component.scss']
})
export class NotebooksComponent implements OnInit {

  constructor(public router: Router, private nb: NotebooksService, private fb: FormBuilder) {
    this.noteBook = this.nb
    this.instanceForm = this.nb.instanceForm
    console.log(this.instanceForm);
    
  }
  ngOnInit(): void {
  }
  noteBook: NotebooksService
  instanceName =''
  instanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    port: ['', Validators.required],
    region: [''],
    zone: [''],
    environment: [''],
    bootDisk: [''],
    subnetwork: [''],
    externalIp: [''],
    permission: [''],
    GPU: [''],
    machineType: ['']
  })
  cancel () {    
    this.noteBook.createInstanceFlag = false
  }
  create () {
    console.log('this.instanceForm', this.instanceForm);
    if (!this.instanceForm.valid) {
      for (const i in this.instanceForm.controls) {
        if (this.instanceForm.controls[i].status === "INVALID") {
          this.instanceName = this.instanceForm.controls[i].value
          console.log('this.instanceForm.get(i)', this.instanceForm.get(i).dirty);
          
        }
      }
      return
    }
    const str = 'HOME.NOTEBOOKS.Create'
    if (this.nb.createInstanceTitle === str) {
      this.nb.createNewApp().subscribe(data => {})
    } else {
      this.nb.modifyApp().subscribe(data => {})
    }
    this.noteBook.createInstanceFlag = false
  }
}
