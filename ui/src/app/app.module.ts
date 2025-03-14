import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { SharedModule } from '@tide-shared/shared.module';
import { HttpClientModule, HttpClient } from '@angular/common/http';
import { TranslateLoader, TranslateModule } from '@ngx-translate/core';
import { TranslateHttpLoader } from '@ngx-translate/http-loader';

import { AppRoutingModule, declarations, providers } from './app-routing.module';
import { AppComponent } from './app.component';
import { I18nService } from '@tide-shared/service/i18n';
import { LandingComponent } from './landing/landing.component';
import { MyLandingFooterComponent } from './layout/my-landing-footer/my-landing-footer.component';
import { VinComponent } from './vin/vin.component';
import { VcppComponent } from './vcpp/vcpp.component';
import { ResetComponent } from './reset/reset.component';

// AoT requires an exported function for factories
export function HttpLoaderFactory(http: HttpClient) {
  return new TranslateHttpLoader(http);
}

@NgModule({
  declarations: [	
    AppComponent,
    ...declarations,
    LandingComponent,
    MyLandingFooterComponent,
    VinComponent,
    VcppComponent,
      ResetComponent
   ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    SharedModule,
    AppRoutingModule,
    HttpClientModule,
    TranslateModule.forRoot({
      defaultLanguage: 'en',
      loader: {
        provide: TranslateLoader,
        useFactory: HttpLoaderFactory,
        deps: [HttpClient],
      },
    }),
  ],
  providers: [
    ...providers,
    I18nService,
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
