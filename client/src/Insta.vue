<template>
  <div id="inin">

  <el-container>
    <el-header>
      <el-col :span="8"><div class="grid-content bg-purple-light"></div></el-col>
      <el-col :span="8"><div class="grid-content bg-purple-light">
        <a href="Top.vue">サイト名</a>
      </div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple-light"></div></el-col>
      <el-col :span="1"><div class="grid-content bg-purple-light">
        <a href="Data.vue"><el-button icon="el-icon-edit" circle></el-button></a>
      </div></el-col>
      <el-col :span="1"><div class="grid-content bg-purple-light">
        <a href="#"><el-button icon="el-icon-search" circle></el-button></a>
      </div></el-col>
    </el-header>
    <el-container>
    <el-aside width="50%">

    </el-aside>
    
    <el-main>
        写真載せるところ
    </el-main>
    </el-container>
    <el-footer>Footer</el-footer>
    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm">
        <el-form-item label="Uer Name" prop="userName">
          <el-input v-model="ruleForm.name"></el-input>
        </el-form-item>
        <el-form-item label="User ID" prop="userId">
          <el-input v-model="ruleForm.id"></el-input>
        </el-form-item>
        <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"
          v-on:vdropzone-sending="sendingEvent"
          v-on:vdropzone-removed-file="removeEvent"
        ></vue-dropzone>
        <el-form-item label="Activity form" prop="desc">
          <el-input type="textarea" v-model="ruleForm.desc"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">Create</el-button>
        </el-form-item>
      </el-form>
  </el-container>



  
  </div>
</template>


<script>
  import axios from 'axios'
  import vue2Dropzone from 'vue2-dropzone'
  import 'vue2-dropzone/dist/vue2Dropzone.min.css'

  export default {
    name: 'Data',
    data() {
      return {
        ruleForm: {
          userName: '',
          userId: '',
          desc: ''
        },
        rules: {
          userName: [
            { required: true, message: 'Please input Activity name', trigger: 'blur' },
            { min: 1, max: 10, message: 'Length should be 1 to 10', trigger: 'blur' }
          ],
          userId: [
            { required: true, message: 'Please input Activity name', trigger: 'blur' },
            { min: 1, max: 10, message: 'Length should be 1 to 10', trigger: 'blur' }
          ],
          desc: [
            { required: true, message: 'Please input activity form', trigger: 'blur' }
          ]
        },
        dropzoneOptions: {
          url: 'http://localhost:8888/images',
          method: 'post',
          addRemoveLinks: 'true'
        }
      };
    },
    components: {
      vueDropzone: vue2Dropzone
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      sendingEvent: function (file, xhr, formData) {
        formData.append('uuid', file.upload.uuid)
      },
      removeEvent: function (file, error, xhr) {
        axios.delete(`http://localhost:8888/images/${file.upload.uuid}`).then(res => {
          console.log(res.data)
        }).catch(err => {
          console.error(err)
        })
      }
    },
    mounted () {
      axios.get('http://localhost:8888/images').then(res => {
        res.data.forEach(res => {
          let filename = res.path.replace('http://localhost:8888/', '')
          let id = filename.replace('.png', '')
          var file = {size: res.size, name: filename, type: "image/png", upload: {uuid: id}}
          this.$refs.myVueDropzone.manuallyAddFile(file, res.path)
        })
      }).catch(err => {
        console.error(err)
      })
    },

  }
</script>




<style>
  .el-header, .el-footer {
    background-color: #B3C0D1;
    color: #333;
    text-align: center;
    line-height: 60px;
  }
  
  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 200px;
  }
  
  .el-main {
    background-color: #E9EEF3;
    color: #333;
    text-align: center;
    line-height: 160px;
  }
  
  body > .el-container {
    margin-bottom: 40px;
  }
  
  .el-container:nth-child(5) .el-aside,
  .el-container:nth-child(6) .el-aside {
    line-height: 260px;
  }
  
  .el-container:nth-child(7) .el-aside {
    line-height: 320px;
  }
</style>