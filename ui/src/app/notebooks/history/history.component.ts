import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit'
import { environment } from '@tide-environments/environment';

@Component({
  selector: 'tide-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {

  constructor(private http:HttpClient) { }
  ngOnInit(): void {
  }
  toGraomacs () {
    const link = environment.webssh   
    let form = document.createElement('form')
    form.action= link
    form.method='get'
    form.target = '_blank'
    const body = document.documentElement
    body.appendChild(form)
    const opts = {
      hostname: '120.133.15.12',
      port: 20023,
      username: 'root',
      password: 'ca$hc0w',
      privatekey: '',
      // passphrase: '',
      totp: ''
    }
    for (const key in opts) {
      const input = document.createElement('input')
      input.type = 'hidden'
      input.name = key
      input.value = opts[key]
      form.appendChild(input)
    }

    setTimeout(() => {
      form.submit()
      form = null
    }, 1000)
  }
}