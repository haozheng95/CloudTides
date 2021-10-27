import { Component, OnInit } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
import { Router } from '@angular/router'
@Component({
  selector: 'tide-create-insrance',
  templateUrl: './create-insrance.component.html',
  styleUrls: ['./create-insrance.component.scss']
})
export class CreateInsranceComponent implements OnInit {
  noteBook: NotebooksService
  instanceForm:NotebookModel = {
    name: '',
    region: '',
    zone: '',
    environment: '',
    machineType: '',
    bootDisk: '',
    subnetwork:'',
    externalIp:'',
    permission: '',
    GPU: ''
  }
  constructor(private noteBooks: NotebooksService, private router: Router) {
    this.noteBook = this.noteBooks
  }

  ngOnInit(): void {
  }
  currentButton = true
  createInstanceFlag = false
  backCreateInstance () {
    this.router.navigate(['/cloudtides/notebooks/list'])
    // this.noteBook.createInstance = true
  }
  openInstanceModal () {
    this.noteBook.createInstanceFlag = true
    this.noteBook.createInstanceTitle = 'HOME.APPLICATION.Create'
  }
  toggleCurrentButton () {
    this.currentButton = !this.currentButton
  }
}
export interface NotebookModel {
  name: string
  region: string
  zone: string
  environment: string
  machineType: string|number
  bootDisk: string
  subnetwork:string
  externalIp:string
  permission: string
  GPU: string
}
