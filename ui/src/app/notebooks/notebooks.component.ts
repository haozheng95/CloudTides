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
  }
  ngOnInit(): void {
  }
  noteBook: NotebooksService
  flag:boolean = false
  errorMsg = ''
  get disabled () {
    if (this.instanceForm.get('instanceName').value !== '' && this.instanceForm.get('port').value !== '') {
      return false
    }
    return true
  }
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
  create (data) {
    console.log('data',data);
    this.errorMsg = ''
    this.flag = true
    const str = 'HOME.APPLICATION.Create'
    const form = {
      instanceName: data.instanceName,
      port: data.port
    }
    if (this.nb.createInstanceTitle === str) {
      this.nb.createNewApp(form).subscribe(data => {
        console.log('data', data);
        this.flag = false
        this.noteBook.createInstanceFlag = false
      },
      err => {
        console.log('err', err);
        
        this.errorMsg = err.statusText
        this.flag = false
      })
    } else {
      this.nb.modifyApp(form).subscribe(data => {
        this.flag = false
        this.noteBook.createInstanceFlag = false
      },
      err => {
        this.errorMsg = err.msg
        this.flag = false
      })
    }
  }
}
