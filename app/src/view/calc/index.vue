<template>
  <div>
    <el-row> </el-row>
    <el-row>
      <el-table :data="tableData" style="width: 100%" border :size='size'>
        <el-table-column type="index" label="index" width="80" align="center" >
        </el-table-column>
        <el-table-column
          prop="description"
          label="description"
          align="center"
          width="180"
        >
        </el-table-column>
        <el-table-column
          prop="expression"
          label="expression"
          show-overflow-tooltip
          align="center"
          width="800"
        >
        </el-table-column>
        <el-table-column
          prop="duration"
          label="duration"
          align="center"
          width="100"
        >
        </el-table-column>
        <el-table-column
          prop="status"
          label="status"
          align="center"
          width="100"
        >
        </el-table-column>
        <el-table-column
          prop="errorMessage"
          label="errorMessage"
          align="center"
          show-overflow-tooltip=""
          width="150"
        >
        </el-table-column>
        <el-table-column
          prop="createTime"
          label="createTime"
          sortable
          align="center"
          width="180"
        >
        </el-table-column>
        <el-table-column
          prop="updatedTime"
          label="updatedTime"
          align="center"
          sortable
          width="180"
        >
        </el-table-column>
        <el-table-column :width="owdith" fixed="right" align="center">
          <template slot="header" slot-scope="scope">
            <el-row>
              <el-col :span="10">
                <el-button type="text" @click="handleCalcOpen" :size="size"  class="hidden-xs-only"
                  >新增计算项</el-button
                >
              </el-col>
              <el-col :span="14">
                <el-input
                class="hidden-xs-only"
                  placeholder="根据description进行搜索"
                  v-model="condition"
                  @input="searchHandler"
                >
                  <i slot="prefix" class="el-input__icon el-icon-search"></i>
                </el-input>
              </el-col>
            </el-row>
          </template>
          <template slot-scope="scope">
            <el-button type="text" :size="size" @click="edit(scope.row)" class="hidden-xs-only">编辑</el-button>
            <el-button
              type="text"
              :size="size"
              v-if="scope.row.status === 'false'"
              @click="startCalc(scope.row)"
              >启动</el-button
            >
            <el-button type="text" v-else :size="size" @click="stopCalc(scope.row)"
              >停止</el-button
            >
            <el-button type="text" :size="size" @click="deleteCalc(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </el-row>
    <!-- 新增计算项 -->
    <el-dialog
      :title="calcDialogName"
      :visible.sync="calcDialog"
      width="800px"
      :showClose="false"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      @opened="handleCalcOpened"
    >
      <el-row>
        <quill-editor
          ref="edit"
          :options="editorOption"
          v-model="calcContent"
        ></quill-editor>
      </el-row>
      <div slot="footer" class="dialog-footer">
        <el-button @click="calcHandler">确定</el-button>
        <el-button @click="calcDialog = false">关闭</el-button>
      </div>
    </el-dialog>
    <!-- 编辑计算项 -->
    <el-dialog
      :title="editDialogName"
      :visible.sync="editDialog"
      width="800px"
      :showClose="false"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      @opened="handleCalcOpened1"
    >
      <el-row>
        <quill-editor
          ref="edit1"
          :options="editorOption1"
          v-model="editContent"
        ></quill-editor>
      </el-row>
      <div slot="footer" class="dialog-footer">
        <el-button @click="updateCalcItem">确定</el-button>
        <el-button @click="editDialog = false">关闭</el-button>
      </div>
    </el-dialog>
    <!-- 代码上传 -->
    <el-dialog
      :title="codeUploadDialogNames"
      :visible.sync="codeUploadDialog"
      width="800px"
      :showClose="false"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-upload
        accept=".js"
        drag
        :action="actionUrl"
        multiple
        :headers="uploadHeaders"
        :http-request="uploadCodeFile"
        :before-upload="beforeCodeUpload"
        :limit="limit"
        :file-list="codeFileList"
        :on-exceed="showExceed"
      >
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip" slot="tip">可以上传js文件</div>
      </el-upload>
      <div slot="footer" class="dialog-footer">
        <el-button @click="showJsCode">确定</el-button>
        <el-button @click="codeUploadDialog = false">关闭</el-button>
      </div>
    </el-dialog>
    <!-- 计算项描述 -->
    <el-dialog
      :title="descriprionDialogNames"
      :visible.sync="descriprionDialog"
      width="550px"
      :showClose="false"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-input v-model="description"  placeholder="请输入计算项描述"></el-input>
      <div slot="footer" class="dialog-footer">
        <el-button @click="descriptionHandler">确定</el-button>
        <el-button @click="descriprionDialog = false">关闭</el-button>
      </div>
    </el-dialog>
    <!-- 编辑时间 -->
    <el-dialog
      :title="timeDurationDialogNames"
      :visible.sync="timeDurationDialog"
      width="550px"
      :showClose="false"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-form label-width="100px">
        <el-form-item
          v-for="(item, index) in editItems"
          :label="item.label"
          :key="index"
        >
          <el-input v-model="item.value"> </el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="timeHandler">确定</el-button>
        <el-button @click="timeDurationDialog = false">关闭</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { getCookie } from "@/utils/cookie";
import { quillEditor } from "vue-quill-editor";
import * as Quill from "quill";
import axios from "axios";
import 'element-ui/lib/theme-chalk/display.css';
import {
  addCalulationItems,
  getCalculationItems,
  startCalcItem,
  updateCalculationItems,
} from "@/api/calc";
export default {
  name: "Calc",
  data: function () {
    return {
      calcDialogName: "二次计算",
      calcDialog: false,
      editorOption: {
        modules: {
          toolbar: {
            container: [
              ["bold", "italic"], //加粗，斜体，下划线，删除线
              ["blockquote", "code-block"], //引用，代码块
              [{ list: "ordered" }, { list: "bullet" }], //列表
              [{ indent: "-1" }, { indent: "+1" }], // 缩进
              [{ size: ["small", false, "large", "huge"] }], // 字体大小
              [{ color: [] }, { background: [] }], // 字体颜色，字体背景颜色
              [{ font: [] }], //字体
              [{ align: [] }], //对齐方式
              ["codeUpload"],
            ],
            handlers: {
              codeUpload: () => {},
            },
          },
        },
        theme: "snow",
      }, // 富文本编辑器
      editorOption1: {
        modules: {
          toolbar: {
            container: [
              ["bold", "italic"], //加粗，斜体，下划线，删除线
              ["blockquote", "code-block"], //引用，代码块
              [{ list: "ordered" }, { list: "bullet" }], //列表
              [{ indent: "-1" }, { indent: "+1" }], // 缩进
              [{ size: ["small", false, "large", "huge"] }], // 字体大小
              [{ color: [] }, { background: [] }], // 字体颜色，字体背景颜色
              [{ font: [] }], //字体
              [{ align: [] }], //对齐方式
              ["editDuration"],
            ],
            handlers: {
              editDuration: () => {},
            },
          },
        },
        theme: "snow",
      }, // 富文本编辑器
      calcContent: "", // 二次计算的输入内容
      codeUploadDialogNames: "代码上传",
      codeUploadDialog: false,
      jsCodeContent: "", // jscode file info
      codeFileList: [], // js code file list
      tableData: [],
      actionUrl: "", // 上传文件的url
      uploadHeaders: { "Content-Type": "multipart/form-data" }, // 上传的头部
      limit: 1, // 文件限制
      descriprionDialogNames: "添加描述",
      descriprionDialog: false,
      description: "", // 描述
      condition: "", // 搜索的条件
      editDialogName: "编辑二次计算项",
      editDialog: false,
      editContent: "", // v-model
      timeDurationDialogNames: "编辑计算项时间",
      timeDurationDialog: false,
      timeDuration: "", // 时间input的v-model
      selectedRow: null, // 选中的行
      editItems: [],
      size: 'medium',
      owdith:'300',
    };
  },
  created() {
    if (document.body.clientWidth < 768){
       this.size = 'mini'
       this.owdith = '150'
    }
    this.actionUrl = "http://" + getCookie("ip") + "/page/uploadFile";
    this.render();
  },
  mounted(){
      document.querySelector(".el-main").style.backgroundColor = " #ffffff";
  },
  methods: {
    handleCalcOpen() {
      this.calcContent = "";
      this.calcDialogName = "新增二次计算项";
      this.calcDialog = true;
    },
    render() {
      getCalculationItems(
        JSON.stringify({
          condition: `description like '%${this.condition}%'`,
        })
      ).then(({ data }) => {
        this.tableData = data.map((item) => {
          if (item.updatedTime === "null") {
            item.updatedTime = "";
          }       
          return item;
        });
      });
    },
    handleCalcOpened() {
      const codeUploadButton = document.querySelector(".ql-codeUpload");
      codeUploadButton.classList = "ql-codeUpload el-button el-button--text";
      codeUploadButton.innerText = "代码上传";
      codeUploadButton.addEventListener("click", () => {
        this.codeUploadDialog = true;
        this.codeFileList = [];
      });
      document.querySelector(".ql-container").style.height = "400px";
    },
    handleCalcOpened1() {
      const editDurationButton = document.querySelector(".ql-editDuration");
      editDurationButton.classList =
        "ql-editDuration el-button el-button--text";
      editDurationButton.innerText = "编辑Item";
      this.editItems = [
        { label: "编辑描述", value: this.selectedRow.description },
        { label: "编辑计算时间", value: this.selectedRow.duration },
      ];
      editDurationButton.addEventListener("click", () => {
        this.timeDurationDialog = true;
      });
      document.querySelector(".ql-container").style.height = "400px";
    },
    calcHandler() {
      if (this.calcContent.length === 0) {
        this.$message.error("计算代码不能为空");
      } else {
        this.descriprionDialog = true;
        this.description = "";
      }
    },
    // 上传js code
    uploadCodeFile() {
      const data = new FormData();
      const fileUps = this.jsCodeContent;
      data.append("file", fileUps);
      axios({
        headers: {
          "Content-Type": "multipart/form-data",
        },
        url: this.actionUrl,
        data: data,
        method: "post",
      })
        .then(() => {
          this.$message.success("上传成功");
        })
        .catch(
          ({
            response: {
              data: { message: message },
            },
          }) => {
            this.$notify.error({
              title: "上传失败",
              message,
            });
          }
        );
    },
    beforeCodeUpload(file) {
      this.jsCodeContent = file;
      if (this.isJs(this.jsCodeContent.name)) {
        return true;
      } else {
        this.$message.error("非法的文件类型或者文件列表已满!");
        return false; // 返回false取消上传
      }
    },
    isJs(name) {
      return /\.(js)$/.test(name);
    },
    showExceed() {
      this.$message.error("文件列表已满");
    },
    showJsCode() {
      axios
        .get(
          "http://" +
            getCookie("ip") +
            "/page/getJsCode/" +
            this.jsCodeContent.name
        )
        .catch(
          ({
            response: {
              data: { message: message },
            },
          }) => {
            this.$notify.error({
              title: "加载二次计算js代码失败",
              message,
            });
          }
        )
        .then(({ data }) => {
          this.calcContent = data.data;
          this.codeUploadDialog = false;
        });
    },
    // 添加二次计算代码
    descriptionHandler() {
      if (this.description.length === 0) {
        this.$message.error("计算项的描述不能为空");
      } else {
        const expression = this.$refs.edit.quill
          .getContents()
          .ops.map((item) => {
            return item.insert;
          })
          .join("");
        addCalulationItems(
          JSON.stringify({
            description: this.description,
            expression,
          })
        )
          .then(() => {
            this.descriprionDialog = false;
            this.calcDialog = false;
            this.$message.success("添加成功");
            this.render();
          })
          .catch(
            ({
              response: {
                data: { message: message },
              },
            }) => {
              this.$notify.error({
                title: "添加失败",
                message,
              });
            }
          );
      }
    },
    // 启动计算
    startCalc({ id }) {
      axios
        .get("http://" + getCookie("ip") + "/calculation/startCalcItem/" + id)
        .then(() => {
          this.$message.success("启动成功!");
          this.render();
        })
        .catch(
          ({
            response: {
              data: { message: message },
            },
          }) => {
            this.$notify.error({
              title: "启动计算项失败",
              message,
            });
          }
        );
    },
    // 终止计算
    stopCalc({ id }) {
      axios
        .get("http://" + getCookie("ip") + "/calculation/stopCalcItem/" + id)
        .then(() => {
          this.$message.success("成功停止计算项");
          this.render();
        })
        .catch(
          ({
            response: {
              data: { message: message },
            },
          }) => {
            this.$notify.error({
              title: "启动计算项失败",
              message,
            });
          }
        );
    },
    // 删除
    deleteCalc({ id }) {
      axios
        .get("http://" + getCookie("ip") + "/calculation/deleteCalcItem/" + id)
        .then(() => {
          this.$message.success("删除成功");
          this.render();
        })
        .catch(
          ({
            response: {
              data: { message: message },
            },
          }) => {
            this.$notify.error({
              title: "删除失败",
              message,
            });
          }
        );
    },
    // 搜索
    searchHandler() {
      this.render();
    },
    // 编辑
    edit(row) {
      this.selectedRow = row;
      this.editContent = row.expression.replace(/\n/g, "<br />");
      this.editDialog = true;
    },
    // 编辑时间确定按钮
    timeHandler() {
      const t = parseFloat(this.editItems[1].value);
      if (60 % t != 0) {
        this.$message.error("计算间隔必须能被60整除");
      } else if (t < 1) {
        this.$message.error("时间间隔不能小于1s");
      } else {
        this.timeDurationDialog = false;
      }
    },
    // 更新二次计算项
    updateCalcItem() {
      const expression = this.$refs.edit1.quill
        .getContents()
        .ops.map((item) => {
          return item.insert;
        })
        .join("");
      const description = this.editItems[0].value;
      const duration = this.editItems[1].value;
      updateCalculationItems(
        JSON.stringify({
          id: this.selectedRow.id,
          description,
          duration,
          expression,
        })
      )
        .then(() => {
          this.editDialog = false;
          this.render();
          this.$message.success("更新成功");
        })
        .catch(
          ({
            response: {
              data: { message: message },
            },
          }) => {
            this.$notify.error({
              title: "更新失败",
              message,
            });
          }
        );
    },
  },
};
</script>