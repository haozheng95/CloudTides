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
    const arr = []
    for (const key in this.instanceForm.controls) {
      if (this.instanceForm.controls[key].value !== 'jupyter' && this.instanceForm.controls[key].value !== 'gromacs') {
        arr.push(this.instanceForm.controls[key].value)
      }
    }
    const rsg = arr.every(el => el !== '')
    return !rsg
  }
  instanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    port: ['', Validators.required],
    appType: [''],
    sshHost: [''],
    sshPassword: [''],
    sshPort: [''],
    sshUser: ['']
  })
  cancel () {    
    this.noteBook.createInstanceFlag = false
    this.instanceForm.setValue({
      instanceName: '',
      port: '',
      appType: '',
      sshHost: '',
      sshPassword: '',
      sshPort: '',
      sshUser: ''
    })
    this.errorMsg = ''
    this.flag = false

  }
  create (data) {
    this.errorMsg = ''
    this.flag = true
    const str = 'HOME.APPLICATION.Create'
    data.sshPort = +data.sshPort
    if (this.nb.createInstanceTitle === str) {
      this.nb.createNewApp(data).subscribe(data => {
        this.flag = false
        this.noteBook.createInstanceFlag = false
        this.instanceForm.setValue({
          instanceName: '',
          port: '',
          appType: '',
          sshHost: '',
          sshPassword: '',
          sshPort: '',
          sshUser: ''
        })
      },
      err => {
        console.log('err', err);
        
        this.errorMsg = err.statusText
        this.flag = false
      })
    } else {
      this.nb.modifyApp(data).subscribe(data => {
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
