import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http'
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class FileService {

  constructor(private httpClient: HttpClient) { }

  getFileList(){
    return this.httpClient.get(environment.gateway+'/file')
  }
  addFile(file:File){
    return this.httpClient.post(environment.gateway+'/file', file);
  }
  deleteFile(file:File){
    return this.httpClient.delete(environment.gateway+'/file'+file.id);
  }
}

export class File{
  id: string;
  name: string;
  content: string;
}
