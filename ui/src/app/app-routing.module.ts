import{RouterModule,Routes} from "@angular/router"
import { HomeComponent } from './home/home.component'
import { FileComponent } from './file/file.component'
import { NgModule } from '@angular/core'
const routes:Routes=[
    {path:'', redirectTo:'home', pathMatch:'full'},
    {path:'home', component: HomeComponent},
    {path:'files', component: FileComponent}
]

@NgModule({
    imports: [ RouterModule.forRoot(routes) ],
    exports: [ RouterModule ]
  })
  export class AppRoutingModule{}