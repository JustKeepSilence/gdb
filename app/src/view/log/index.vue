<template>
  <div>
    <el-row>
      <el-col :span="5">
        <el-form label-width="80px">
          <el-form-item label="日志筛选">
            <el-select
              v-model="selecteType"
              placeholder="请选择"
              @change="changeLogs"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="5">
        <el-radio-group v-model="rev">
          <el-radio :label="1" border>正序</el-radio>
          <el-radio :label="2" border @change="revChange">倒序</el-radio>
        </el-radio-group>
      </el-col>
    </el-row>
    <el-row>
      <el-timeline>
        <el-timeline-item
          v-for="(item, index) in logs"
          :key="index"
          placement="top"
          :timestamp="item.time"
        >
          <el-card>
            <el-row
              ><p style="float: left">Level: {{ item.level }}</p></el-row
            >
            <el-row
              ><p style="float: left">
                RequestMethod:{{ item.method }}
              </p></el-row
            >
            <el-row
              ><p style="float: left">
                RequestString:{{ item.requestString }}
              </p></el-row
            >
            <el-row
              ><p style="float: left">RequestUrl:{{ item.url }}</p></el-row
            >
            <el-row
              ><p style="float: left">LogMessage:{{ item.msg }}</p></el-row
            >
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </el-row>
  </div>
</template>
<script>
import axios from "axios";
import { getCookie } from "@/utils/cookie";
export default {
  name: "Log",
  data() {
    return {
      logs: [],
      options: [
        { label: "All", value: "all" },
        { label: "Info", value: "info" },
        { label: "Error", value: "error" },
      ],
      selecteType: "all",
      rev: 1,
    };
  },
  watch:{
      rev(n,o){
           this.rev = n
          this.getLogs()
      }
  },
  mounted(){
      document.querySelector(".el-main").style.backgroundColor = " #ffffff";
  },
  created() {
    this.getLogs();
  },
  methods: {
    changeLogs() {
      this.getLogs();
    },
    getLogs() {
      let d = null;
      axios
        .get("http://" + getCookie("ip") + "/page/getLogs")
        .then(({ data: { data } }) => {
          if (this.selecteType === "all") {
            d = data;
          } else {
            d = data.filter((item) => {
              return item.level === this.selecteType;
            });
          }
          if (this.rev == 1) {
            this.logs = d.sort((a, b) => {
              return new Date(a.time).getTime() > new Date(b.time).getTime()
                ? 1
                : -1;
            });
          } else {
            this.logs = d.sort((a, b) => {
              return new Date(a.time).getTime() > new Date(b.time).getTime()
                ? -1
                : 1;
            });
          }
        })
        .catch(({ response: { data: message } }) => {
          this.$notify({
            title: "获取日志失败",
            message,
          });
        });
    },
  },
};
</script>
<style scoped>
</style>