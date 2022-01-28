import { ModuleWithProviders, NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { ClarityModule } from '@clr/angular';
import { DropzoneDirective } from './directive/dropzone.directive';
import { MessageModule } from './component/message/message.module'

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    MessageModule,
    ReactiveFormsModule,
    HttpClientModule,
    ClarityModule,
  ],
  exports: [
    CommonModule,
    MessageModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    ClarityModule,
    DropzoneDirective
  ],
  declarations: [
    DropzoneDirective
  ],
})
export class BaseModule {

  static forRoot(): ModuleWithProviders {
    return {
      ngModule: BaseModule,
      providers: [],
    };
  }

  static forChild(): ModuleWithProviders {
    return {
      ngModule: BaseModule,
      providers: [],
    };
  }
}

