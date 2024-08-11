<script>
import router from "@/router";
import {CreateConfig, DeleteConfig, GetAllConfig, LoadConfig, UpdateConfig} from "@/view/api/conf";
import {BackUpAll, RecoverBackUp} from "@/view/api/tool";

export default {
  data() {
    return {
      activeName: "first",
      configs: [],
      selectVersion:'',
      showAddConfig: false,
      showBackUpUpload:false,
      addConfigForm: {
        version: "",
        open: true,
        upload: {
          path: "",
          domain: "",
          maxsize: "4MB",
          bingurl: "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
        },
        backup: {
          cycle: 7,
          openemail: false
        },
        qqemail: {
          username: "",
          password: "",
          host: "smtp.qq.com",
          port: 465
        },
        admin: {
          username: "",
          password: ""
        }
      }
    };
  },
  methods: {
    RecoverBackUp,
    BackUpAll,
    logout() {
      localStorage.clear();
      router.go(0);
    },
    getAllConfig() {
      GetAllConfig().then(res => {
        this.configs = res.data.map(config => ({ ...config, editing: false }));
      });
    },
    enableEditing(configKey) {
      this.configs[configKey].editing = true;
    },
    updateConfig(configKey) {
      this.configs[configKey].editing = false;
      const data = this.configs[configKey];
      UpdateConfig(data).then(() => {
        this.$message.success("更新成功")
        this.getAllConfig();
      })
    },
    saveConfig() {
      if (
          this.addConfigForm.admin.username === '' ||
          this.addConfigForm.admin.password === ''
      ) {
        this.$message.warning("请填写所有必填字段");
        return;
      }
      this.$prompt("请输入版本号", {
        inputValue: '',
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).then(({ value }) => {
        if (value === '') {
          this.$message.warning("版本号不能为空")
          return;
        }
        this.addConfigForm.version = value;
        CreateConfig(this.addConfigForm).then(() => {
          this.$message.success("添加成功")
          this.getAllConfig();
        })
      }).catch(() => {
        this.$message.info("取消")
      })
    },
    deleteConfig() {
      if (this.selectVersion === '') {
        this.$message.warning("请选择配置")
      } else {
        this.$confirm('确定删除该配置吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // Implement logic to delete the selected config from the backend
          DeleteConfig(this.selectVersion).then(() => {
            this.$message.success('删除成功');
            this.getAllConfig();
          })
        }).catch(() => {
          this.$message.info('已取消删除');
        });
      }
    },
    loadConfig() {
      // Implement logic to load the selected config from the backend
      if (this.selectVersion === '') {
        this.$message.warning("请选择配置")
      } else {
        this.$confirm('确定加载该配置吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // Implement logic to delete the selected config from the backend
          LoadConfig(this.selectVersion).then(() => {
            this.$message.success('加载成功');
            this.getAllConfig();
          })
        }).catch(() => {
          this.$message.info('已取消加载');
        });
      }
    },
    cancelEditing(configKey) {
      this.configs[configKey].editing = false;
    }
  },
  mounted() {
    this.getAllConfig();
  }
};
</script>

<template>
  <el-container>
    <el-main>
      <el-tabs v-model="activeName" class="demo-tabs">
        <el-tab-pane label="系统设置" name="first">
          <el-dialog v-model="showAddConfig" title="添加配置" width="40%" style="border-radius: 20px;">
            <el-collapse>
              <el-collapse-item title="上传配置">
                <el-descriptions border column="1">
                  <el-descriptions-item label="路径">
                    <el-input v-model="addConfigForm.upload.path" ></el-input>
                  </el-descriptions-item>
                  <el-descriptions-item label="域">
                    <el-input v-model="addConfigForm.upload.domain"></el-input>
                  </el-descriptions-item>
                  <el-descriptions-item label="最大内存">
                    <el-input v-model="addConfigForm.upload.maxsize" ></el-input>
                  </el-descriptions-item>
                  <el-descriptions-item label="Bing URL">
                    <el-input v-model="addConfigForm.upload.bingurl"></el-input>
                  </el-descriptions-item>
                </el-descriptions>
              </el-collapse-item>
              <el-collapse-item title="备份配置">
                <el-descriptions border column="1">
                  <el-descriptions-item label="周期">
                    <el-input-number v-model="addConfigForm.backup.cycle" :min="0"></el-input-number>
                  </el-descriptions-item>
                  <el-descriptions-item label="开启邮件">
                    <el-select v-model="addConfigForm.backup.openemail">
                      <el-option label="是" :value="true"></el-option>
                      <el-option label="否" :value="false"></el-option>
                    </el-select>
                  </el-descriptions-item>
                </el-descriptions>
              </el-collapse-item>
              <el-collapse-item title="邮箱配置">
                <el-descriptions border column="1">
                  <el-descriptions-item label="用户名">
                    <el-input v-model="addConfigForm.qqemail.username"></el-input>
                  </el-descriptions-item>
                  <el-descriptions-item label="密码">
                    <el-input v-model="addConfigForm.qqemail.password" ></el-input>
                  </el-descriptions-item>
                  <el-descriptions-item label="主机">
                    <el-input v-model="addConfigForm.qqemail.host"></el-input>
                  </el-descriptions-item>
                  <el-descriptions-item label="端口">
                    <el-input v-model="addConfigForm.qqemail.port"></el-input>
                  </el-descriptions-item>
                </el-descriptions>
              </el-collapse-item>
              <el-collapse-item title="管理员账户配置">
                <el-descriptions border column="1">
                  <el-descriptions-item label="用户名">
                    <el-input v-model="addConfigForm.admin.username"></el-input>
                  </el-descriptions-item>
                  <el-descriptions-item label="密码">
                    <el-input v-model="addConfigForm.admin.password"></el-input>
                  </el-descriptions-item>
                </el-descriptions>
              </el-collapse-item>
            </el-collapse><br>
            <el-button type="primary" @click="saveConfig">保存</el-button>
            <el-button @click="showAddConfig = false">取消</el-button>
          </el-dialog>
          <div class="config-cards">
            <div style="display: flex;">
              <el-select
                  v-model="selectVersion"
                  placeholder="Select"
                  style="margin-right: 10px;"
              >
                <el-option
                    v-for="item in configs"
                    :key="item.Version"
                    :label="item.Version"
                    :value="item.Version"
                />
              </el-select>
              <el-button type="primary" @click="showAddConfig = true" style="margin-bottom: 20px;">添加配置</el-button>
              <el-button type="info" @click="loadConfig">加载配置</el-button>
              <el-button type="danger" @click="deleteConfig">删除配置</el-button>
            </div>
            <el-row :gutter="20">
              <el-col :span="8" v-for="(config, configKey) in configs" :key="config.Version">
                <el-card class="box-card">
                  <div>
                    <span>版本:{{ config.Version }}
                    </span>
                    <el-tag :type="config.Open ? 'success' : 'danger'" class="config-status">
                      {{ config.Open ? '启用' : '暂存' }}
                    </el-tag>
                  </div><br>
                  <el-collapse>
                    <el-collapse-item title="上传配置">
                      <el-descriptions border column="1">
                        <el-descriptions-item label="路径">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.Upload.Path }}
                          </span>
                          <el-input v-else v-model="config.Upload.Path"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="域">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.Upload.Domain }}
                          </span>
                          <el-input v-else v-model="config.Upload.Domain"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="最大内存">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.Upload.MaxSize }}
                          </span>
                          <el-input v-else v-model="config.Upload.MaxSize"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="Bing URL">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.Upload.BingUrl }}
                          </span>
                          <el-input v-else v-model="config.Upload.BingUrl"></el-input>
                        </el-descriptions-item>
                      </el-descriptions>
                    </el-collapse-item>
                    <el-collapse-item title="备份配置">
                      <el-descriptions border column="1">
                        <el-descriptions-item label="周期">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.BackUp.Cycle }}
                          </span>
                          <el-input v-else v-model="config.BackUp.Cycle"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="开启邮件">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.BackUp.OpenEmail ? '是' : '否' }}
                          </span>
                          <el-select v-else v-model="config.BackUp.OpenEmail">
                            <el-option label="是" :value="true"></el-option>
                            <el-option label="否" :value="false"></el-option>
                          </el-select>
                        </el-descriptions-item>
                      </el-descriptions>
                    </el-collapse-item>
                    <el-collapse-item title="邮箱配置">
                      <el-descriptions border column="1">
                        <el-descriptions-item label="用户名">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.QQEmail.Username }}
                          </span>
                          <el-input v-else v-model="config.QQEmail.Username"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="密码">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.QQEmail.Password }}
                          </span>
                          <el-input v-else v-model="config.QQEmail.Password"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="主机">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.QQEmail.Host }}
                          </span>
                          <el-input v-else v-model="config.QQEmail.Host"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="端口">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.QQEmail.Port }}
                          </span>
                          <el-input v-else v-model="config.QQEmail.Port"></el-input>
                        </el-descriptions-item>
                      </el-descriptions>
                    </el-collapse-item>
                    <el-collapse-item title="管理员账户配置">
                      <el-descriptions border column="1">
                        <el-descriptions-item label="用户名">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.Admin.Username }}
                          </span>
                          <el-input v-else v-model="config.Admin.Username"></el-input>
                        </el-descriptions-item>
                        <el-descriptions-item label="密码">
                          <span v-if="!config.editing" @dblclick="enableEditing(configKey)">
                            {{ config.Admin.Password }}
                          </span>
                          <el-input v-else v-model="config.Admin.Password"></el-input>
                        </el-descriptions-item>
                      </el-descriptions>
                    </el-collapse-item>
                  </el-collapse>
                  <el-button v-if="config.editing" type="primary" @click="updateConfig(configKey)">保存</el-button>
                  <el-button v-if="config.editing" @click="cancelEditing(configKey)">取消</el-button>
                </el-card>
              </el-col>
            </el-row>
          </div>
        </el-tab-pane>
        <el-tab-pane label="备份管理" name="second">
          <el-button type="info" @click="BackUpAll()">导出备份</el-button>
          <el-button type="success" @click="showBackUpUpload = true">备份还原</el-button>
          <el-dialog title="备份还原" style="border-radius: 20px;" v-model="showBackUpUpload" width="40%">
            <el-upload
                class="upload-demo"
                drag
                :http-request="RecoverBackUp"
                multiple
            >
              <div class="el-upload__text">
                拖拽文件到此处，或<em>点击上传</em>
              </div>
              <template #tip>
                <div class="el-upload__tip">
                  仅限 zip 文件
                </div>
              </template>
            </el-upload>
          </el-dialog>
        </el-tab-pane>
        <el-tab-pane label="账户设置" name="third">
          <el-form>
            <el-form-item>
              <el-button type="danger" @click="logout" >退出登录</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<style scoped>
.config-cards {
  padding: 20px;
}

.config-status {
  float: right;
}

.el-card {
  margin-bottom: 20px;
}

.el-collapse-item {
  margin-bottom: 10px;
}

.box-card {
  border-radius: 20px;
  box-shadow: 10px 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.delete-button {
  margin-top: 10px;
}
@import url('../../assets/css/main.css');
</style>
