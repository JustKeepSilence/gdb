<template>
  <div style="backgound-color:#909399;width:100%;height:100%">
  <div class="container" v-if="!small">
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" id="loginForm">
      <el-form-item prop="ip">
         <el-input
          v-model="ruleForm.ip"
          autocomplete="off"
          prefix-icon="el-icon-position"
          placeholder="请输入远端服务器地址"
        ></el-input>
      </el-form-item>
      <el-form-item prop="userName">
        <el-input
          v-model="ruleForm.userName"
          autocomplete="off"
          prefix-icon="el-icon-user-solid"
          placeholder="请输入用户名"
        ></el-input>
      </el-form-item>
      <el-form-item prop="passWord">
        <el-input
          type="password"
          v-model="ruleForm.passWord"
          autocomplete="off"
          prefix-icon="el-icon-key"
          placeholder="请输入密码"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm" style="width: 100%">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="mobile_container" v-else>
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" id="loginForm">
      <el-form-item prop="ip">
         <el-input
          v-model="ruleForm.ip"
          autocomplete="off"
          prefix-icon="el-icon-position"
          placeholder="请输入远端服务器地址"
        ></el-input>
      </el-form-item>
      <el-form-item prop="userName">
        <el-input
          v-model="ruleForm.userName"
          autocomplete="off"
          prefix-icon="el-icon-user-solid"
          placeholder="请输入用户名"
        ></el-input>
      </el-form-item>
      <el-form-item prop="passWord">
        <el-input
          type="password"
          v-model="ruleForm.passWord"
          autocomplete="off"
          prefix-icon="el-icon-key"
          placeholder="请输入密码"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm" style="width: 100%">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
  </div>
</template>

<script>
import { userLogin, passWordValidator,readTextFile } from "@/api/login";
import { setCookie } from '@/utils/cookie'
import axios from 'axios'
export default {
  name: "Login",
  data() {
    return {
      ruleForm: {
        userName: "",  // 用户名
        passWord: "",  // 密码
        ip: ""
      },
      
      rules: {
        userName: [{ required: true, trigger: "blur", message: '用户名不能为空' }],  // 用户名的验证
        passWord: [
          {
            required: true,
            tigger: "blur",
            min: 6,
            validator: passWordValidator,  // 自定义验证函数
          },
        ],
      },
      userName : '',
      small: false
    }
  },
  mounted(){
    if (document.body.clientWidth < 768){
       this.small = true
    }
  },
  methods: {
    submitForm() {
      this.userName = this.ruleForm.userName
      setCookie({key: "ip", value: this.ruleForm.ip})  // set ip
      userLogin(
        JSON.stringify({
          username: this.ruleForm.userName,
          password: this.ruleForm.passWord,
        })
      ).then(({data}) => {
        if (data === null){
        // 当登陆成功以后将token写入cookie,token即为username
        this.$store.dispatch("user/setToken", this.userName).then(() => {
          this.$router.push("/index"); // 跳转到首页
        })}else{
          this.$message.error(data)
        }
      }).catch((e)=>{
          this.$message.error(e.message)
      })
    }
  },
};
</script>
<style scoped>
.container {
  width: 400px;
  margin-left: 600px;
  margin-top: 370px;
}
.mobile_container{
  width: 300px;
  margin-top: 200px;
}
.background{
    width:100%;  
    height:100%;  /**宽高100%是为了图片铺满屏幕 */
    z-index:-1;
    position: absolute;
}
.front{
    z-index:1;
    position: absolute;
}
</style>