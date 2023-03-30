<script lang="ts" setup>
import {reactive} from 'vue'
import {Greet, SelectFolder, ArchApp} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Select Destination Folder and Click to Generate Logs 👇",
  pathSelected: "",
  class: "result0",
  arch: "",
  disabledPath: false,
  disableBtnGenerateLogs: true
})

function awaiting() {
  if (data.arch == "windows") {
    data.resultText = "Gathering System Information - Please allow a few minutes"
  }else if(data.arch == "macos"){
    data.resultText = "Gathering System Information - Please allow a few minutes"
  }else{
    data.resultText = "Gathering System Information - Please allow a few minutes"
  }
    
  data.class = "result1"
}

function disableBtn(event:any) {
  if (data.arch == "windows") {
    data.resultText = "Gathering System Information - Please allow a few minutes"
  }else if(data.arch == "macos"){
    data.resultText = "Gathering System Information - Please allow a few minutes"
  }else{
    data.resultText = "Gathering System Information - Please allow a few minutes"
  }

  event.target.disabled = true
  data.class = "result1"
  data.disabledPath = true
}


function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
    data.class = "result2"
  })
}

function selectDir(){
  SelectFolder().then(
    res => {
      console.log(res);
      data.name = res;
      getArch();
      if(res != ""){
        data.disableBtnGenerateLogs = false;
      }
    }
  )
}

function checkInputDir(event:any){
  if(event.target.value != ""){
    data.disableBtnGenerateLogs = false;
  }else{
    data.disableBtnGenerateLogs = true;
  }
}

function getArch(){
  ArchApp().then(res => {data.arch = res})
}

</script>

<template>
  <main>
    <div id="result" v-bind:class="data.class">{{ data.resultText }}</div>
    <div class="ui large action input" style="width:-webkit-fill-available;margin: 3% 10%;">
      <input @change="checkInputDir($event)" class="input" id="pathFile" type="text" v-model="data.name" placeholder="Select path to output reports..." name="pathFile"
        :disabled="data.disabledPath">
      <button class="ui icon inverted violet button" @click="selectDir()" :disabled="data.disabledPath">
        <i class="folder open icon"></i>
      </button>
    </div>

    <div class="input-box">
        <button id="btn1" class="ui inverted violet button" @click="awaiting();disableBtn($event);greet()" :disabled="data.disableBtnGenerateLogs">Generate Log Files</button>
    </div>
  </main>
</template>

<style scoped>
.result0 {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  color: white;
}
.result1 {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  color: rgb(243, 0, 0);
}
.result2 {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  color: rgb(0, 215, 4);
}

.input-box .btn {
  cursor: pointer;
  display: inline-block;
  min-height: 1em;
  outline: 0;
  border: none;
  vertical-align: baseline;
  /* background: #e0e1e2 none; */
  color: rgba(62,62,116,.6);
  font-family: Lato,'Helvetica Neue',Arial,Helvetica,sans-serif;
  margin: 0 .25em 0 0;
    margin-bottom: 0px;
  padding: .78571429em 1.5em .78571429em;
  text-transform: none;
  text-shadow: none;
  font-weight: 700;
  line-height: 1em;
  font-style: normal;
  text-align: center;
  text-decoration: none;
  border-radius: .28571429rem;
  box-shadow: 0 0 0 1px transparent inset,0 0 0 0 rgba(34,36,38,.15) inset;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  transition: opacity .1s ease,background-color .1s ease,color .1s ease,box-shadow .1s ease,background .1s ease;
  will-change: '';
  -webkit-tap-highlight-color: transparent;
}


.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
