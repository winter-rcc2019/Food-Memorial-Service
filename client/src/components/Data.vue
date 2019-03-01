<template>
  <div class="data">

  <el-container>
    <el-header>サイト名</el-header>
    <center>投稿画面</center>
    <el-row>
      <el-col :span="8"><div class="grid-content bg-purple-light"></div></el-col>
      <el-col :span="8"><div class="grid-content bg-purple-light">
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="10em" class="demo-ruleForm">
          <el-form-item label="ユーザー名" prop="userName">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>
          <el-form-item label="ユーザーID" prop="userId">
            <el-input v-model="ruleForm.id"></el-input>
          </el-form-item>
          <el-form-item label="画像アップロード">
            <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"
              v-on:vdropzone-sending="sendingEvent"
              v-on:vdropzone-removed-file="removeEvent"
            ></vue-dropzone>
          </el-form-item>
          <el-form-item label="概要" prop="desc">
            <el-input type="textarea" v-model="ruleForm.desc"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">投稿</el-button>
          </el-form-item>
        </el-form>
      </div></el-col>
      <el-col :span="8"><div class="grid-content bg-purple-light"></div></el-col>
    </el-row>
    <el-footer>Footer</el-footer>
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
            { required: true, message: 'Please input Activity name', trigger: 'blur' }
          ],
          userId: [
            { required: true, message: 'Please input Activity name', trigger: 'blur' }
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
        formData.append('userName', ruleForm.userName)
        formData.append('userId', ruleForm.userName)
        formData.append('uuid', file.upload.uuid)
        formData.append('desc', ruleForm.desc)
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
  .el-header {
    margin-bottom:  3em;
    background-color: #B3C0D1;
    color: #333;
    text-align: center;
    line-height: 60px;
  }

    .el-footer {
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

  .el-form-item{
    margin: 20px 10px
  }

  .el-row {
    margin: 20px 0;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    
  }

  .el-button{
    margin: 30px auto 10px auto;

  }

  .bg-purple-light {
    /*background: #e5e9f2;*/
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }

  #app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 60px 20px;
}
</style>