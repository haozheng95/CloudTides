import { Component, OnInit, OnDestroy } from '@angular/core';
import { EMPTY, Subject } from 'rxjs';

import { LoginService } from './login.service';
import { catchError, switchMap, tap } from 'rxjs/operators';
import { Router } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
import { I18nService } from '@tide-shared/service/i18n';
// import { Session } from 'inspector';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'cp-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit, OnDestroy {

  constructor(
    public readonly loginService: LoginService,
    private readonly router: Router,
    public readonly translate: TranslateService,
    public readonly i18nService: I18nService,
  ) {}

  readonly vo = {

    model: {
      username: '',
      password: '',
    },

    submitting: false,
    loginError: '',
  };

  readonly session = new BehaviorSubject<UserInfo>({} as any);

  private readonly submit$ = new Subject<Credential>();

  private readonly submit$$ = this.submit$.asObservable()
    .pipe(
      tap(() => {
        this.vo.submitting = true;
        this.vo.loginError = '';
      }),
      switchMap(({ username, password }) => {

        return this.loginService
          .login(username.trim(), password)
          .pipe(
            tap(() => {
              this.vo.submitting = false;
            }),
            catchError((error, source) => {
              this.vo.submitting = false;
              this.vo.loginError = error.message;

              return EMPTY as typeof source;
            }),
          );
      }),
    )
    .subscribe(res => {
      // this.document.location.href = '/'
      
      console.log("here")
      console.log("222:"+localStorage.getItem("pwReset"))
      
      this.router.navigate(['/cloudtides']);
      console.log("333:"+localStorage.getItem("pwReset"))
      if (localStorage.getItem("pwReset") === "false") {
        console.log("entered!!!")
        this.router.navigate(['/cloudtides/reset']);
        console.log("111111")
        // this.router.navigate(['/cloudtides/resource']);
      }

    })
  ;

  onSubmit({ username = '', password = '' }: Credential) {
    this.submit$.next({ username, password });
  }

  ngOnInit() {}

  ngOnDestroy() {
    this.submit$$.unsubscribe();
  }
}

interface Credential {
  username: string;
  password: string;
}

export interface UserInfo {
  username: string;
  priority: string;
  firstName: string;
  lastName: string;
  country: string;
  city: string;
  companyName: string;
  position: boolean;
  email: string,
  phone: string,
  role: string,
  pwReset: boolean,
  orgName: string,
}
