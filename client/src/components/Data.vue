<template>
  <div class="data">
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
      
      <center>投稿画面</center>
      <el-row>
        <el-col :span="8"><div class="grid-content bg-purple-light"></div></el-col>
        <el-col :span="8"><div class="grid-content bg-purple-light">
          <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="10em" class="demo-ruleForm">
            <el-form-item label="ユーザー名" prop="user_name">
              <el-input v-model="ruleForm.user_name"></el-input>
            </el-form-item>
            <el-form-item label="ユーザーID" prop="user_id">
              <el-input v-model="ruleForm.user_id"></el-input>
            </el-form-item>
            <el-form-item label="画像アップロード">
              <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"
                v-on:vdropzone-sending="sendingEvent"
                v-on:vdropzone-removed-file="removeEvent"
              ></vue-dropzone>
            </el-form-item>
            <el-form-item label="概要" prop="comment">
              <el-input type="textarea" v-model="ruleForm.comment"></el-input>
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
          user_name: '',
          user_id: '',
          comment: ''
        },
        rules: {
          user_name: [
            { required: true, message: 'Please input Activity name', trigger: 'blur' }
          ],
          user_id: [
            { required: true, message: 'Please input Activity name', trigger: 'blur' }
          ],
          comment: [
            { required: true, message: 'Please input activity form', trigger: 'blur' }
          ]
        },
        dropzoneOptions: {
          url: 'http://localhost:8888/images/photo',
          method: 'post',
          addRemoveLinks: 'true'
        }
      };
    },
    components: {
      vueDropzone: vue2Dropzone
    },
    methods: {
      submitForm(form_name) {
        this.$refs[form_name].validate((valid) => {
          if (valid) {
            alert('submit!');

          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(form_name) {
        this.$refs[form_name].resetFields();
      },
      sendingEvent: function (file, xhr, formData) {
        formData.append('user_name', ruleForm.user_name)
        formData.append('user_id', ruleForm.user_id)
        formData.append('uuid', file.upload.uuid)
        formData.append('comment', ruleForm.comment)
      },
      removeEvent: function (file, error, xhr) {
        axios.delete(`http://localhost:8888/images/photo/${file.upload.uuid}`).then(res => {
          console.log(res.data)
        }).catch(err => {
          console.error(err)
        })
      }
    },
    mounted () {
      axios.get('http://localhost:8888/images/photo').then(res => {
        res.data.forEach(res => {
          let filename = res.path.replace('http://localhost:8888/', '')
          let id = filename.replace('.png', '')
          var file = {size: res.size, name: filename, type: "image/photo/png", upload: {uuid: id}}
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

  .el-button{
    margin: auto 10px auto;
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