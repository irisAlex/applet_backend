<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="部门" style="width:200px;" prop="method">
                    <el-select v-model="searchInfo.department" placeholder="选择部门">
                        <el-option v-for="item in departmentList" :key="item.authorityId" :label="item.authorityName"
                            :value="item.authorityName">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="类型" style="width:200px;" prop="method">
                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="类别" prop="method" style="width:200px;">
                    <el-select v-model="searchInfo.category" placeholder="请选择">
                        <el-option v-for="item in genreList1" :key="item.name" :label="item.genre" :value="item.genre">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="受检物名称" prop="method" style="width:200px;">
                    <el-input v-model="searchInfo.checkout_name" placeholder="受检物名称" />
                </el-form-item>
                <el-form-item label="供应商名称" prop="method" style="width:200px;">
                    <el-input v-model="searchInfo.supplier" placeholder="供应商名称" />
                </el-form-item>
                <el-form-item label="项目名称" prop="project" style="width:200px;">
                    <el-input v-model="searchInfo.project" placeholder="项目名称" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog('setting')">新增</el-button>
            </div>
            <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
                <el-table-column align="left" label="ID" min-width="150" prop="ID" />
                <el-table-column align="left" label="编号" min-width="150" prop="serialnumber" />
                <el-table-column align="left" label="部门" min-width="150" prop="department" />
                <el-table-column align="left" label="类型" min-width="150" prop="mold" />
                <el-table-column align="left" label="类别" min-width="150" prop="category" />
                <el-table-column align="left" label="受检物名称" min-width="150" prop="checkout_name" />
                <el-table-column align="left" label="受检物号" min-width="150" prop="checkout_number" />
                <el-table-column align="left" label="处理方式" min-width="150" prop="process_mode" />
                <el-table-column align="left" label="处理过程" min-width="150" prop="process_mode">
                    <template #default="scope">
                        <el-progress :text-inside="true" :stroke-width="26"
                            :percentage="scope.row.a3_step === 9 ? 100 : (scope.row.a3_step * 10)"
                            status="success"></el-progress>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="当前负责人" min-width="150" prop="" sortable="custom">
                    <template #default="scope">
                        <el-tag type="info">{{ findPrincipal(scope.row, scope.row.a3_step) }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="计划时间" min-width="150" prop="checkout_date" sortable="custom">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.checkout_date) }}</span>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="填表日期" min-width="150" prop="created_at" sortable="custom">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.created_at) }}</span>
                    </template>
                </el-table-column>

                <el-table-column align="left" fixed="right" label="操作" width="300">
                    <template #default="scope">
                        <el-button icon="download" type="primary" link @click="downExcel(scope.row)">下载</el-button>
                        <el-button type="primary" icon="setting" link @click="openDialog(scope.row, 'setting')">
                            配置
                        </el-button>
                        <el-button icon="edit" type="primary" link
                            @click="openDialog(scope.row, 'report')">填写报告</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>

        </div>
        <!-- issue 确认 -->
        <el-drawer v-model="drawer" :title="dialogTitle" :direction="direction" :before-close="handleClose" :size="ize">
            <div class="block">
                <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                    <el-timeline>
                        <el-timeline-item timestamp="考情分析" placement="top">
                            <el-card>
                                <el-form-item label=" 招考人数:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>

                                <el-form-item label="入围比例:" prop="station" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>

                                <el-form-item label="招考年份:" prop="checkout_name" style="width:50%">
                                    <el-date-picker v-model="value3" type="year" placeholder="选择年">
                                    </el-date-picker>
                                </el-form-item>

                                <el-form-item label="进面分数线:" prop="serialnumber" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>

                                <el-form-item label="收据来源:" prop="defect_problem" style="width:50%">
                                    <span>{{ form.defect_problem }}</span>
                                </el-form-item>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="岗位信息" placement="top">
                            <el-card>
                                <el-form-item label="岗位名称:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="岗位代码:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="岗位类别:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="从事工作:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="单位名称:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="单位序号:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="报考条件" placement="top">
                            <el-card>
                                <el-form-item label="来源类别:" prop="area" style="width:50%">
                                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="学历要求:" prop="area" style="width:50%">
                                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="专业科目:" prop="area" style="width:50%">
                                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="所学专业:" prop="area" style="width:50%">
                                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="专业名称:" prop="area" style="width:50%">
                                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="学位要求:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="职称要求:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="职业资格:" prop="area" style="width:50%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="其他条件:" prop="area" style="width:50%;">
                                    <el-input type="textarea" placeholder="请输入内容" v-model="form.craft_view"
                                        maxlength="50" show-word-limit :rows="10" />
                                </el-form-item>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp=" 上岸课程" placement="top">
                            <el-card>
                                视频
                            </el-card>
                        </el-timeline-item>
                    </el-timeline>
                </el-form>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="handleClose">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-drawer>
    </div>
</template>
<script setup>
import {
    createSupplierApi,
    getSupplierList,
    getSupplierById,
    deleteSupplier,
    updateSupplier

} from '@/api/supplier'
import { toSQLLine } from '@/utils/stringFun'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
const quillEditor = ref();

const editorOptions = reactive({
    modules: {
        toolbar: [  // 工具栏配置
            ['bold', 'italic', 'underline', 'strike'],  // 粗体、斜体、下划线、删除线
            [{ 'header': 1 }, { 'header': 2 }],  // 标题1和标题2
            [{ 'list': 'ordered' }, { 'list': 'bullet' }],  // 有序列表和无序列表
            [{ 'script': 'sub' }, { 'script': 'super' }],  // 上标和下标
            [{ 'indent': '-1' }, { 'indent': '+1' }],  // 缩进
            [{ 'direction': 'rtl' }],  // 文字方向
            [{ 'size': ['small', false, 'large', 'huge'] }],  // 字号
            [{ 'header': [1, 2, 3, 4, 5, 6, false] }],  // 标题等级
            [{ 'color': [] }, { 'background': [] }],  // 字体颜色和背景色
            [{ 'font': [] }],  // 字体
            [{ 'align': [] }],  // 对齐方式
            ['clean']  // 清除格式
        ]
    }
})

defineOptions({
    name: 'Setting',
})


const dialogFormVisible = ref(false)
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const handlePictureCardPreview = (file) => {
    dialogImageUrl.value = file.url;
    dialogVisible.value = true;
    dialogFormVisible.value = true
}

const drawer = ref(false)
const direction = ref('rtl')
const ize = ref('40%')
const dialogTitle = ref('岗位信息配置')
const openDialog = (row, key) => {
    switch (key) {
        case 'setting':
            dialogTitle.value = '岗位信息配置'
            break
        case 'modify':
            dialogTitle.value = '岗位信息修改'
            showSettingDrawer(row)
            break
        default:
            break
    }
    drawer.value = true
}

const handleClose = () => {
    initForm()
    initDraw()
    drawer.value = false
}
const ID = ref(1)
const showSettingDrawer = async (row) => {
    const res = await getManageById({ id: row.ID })
    ID.value = row.ID
    if (res.code === 0) {
        form.value = res.data.manage
    }
}
const setp = ref(0)
const isPass = ref(0)
const imgBase64 = ref('')
const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'

const handleChange = (ste) => {
    setp.value = ste
}
const initDraw = () => {
}

const apis = ref([])
const departmentList = ref([])
const genreList1 = ref([])
const supplierList = ref([])
const projectList = ref([])
const fileList = ref([])
const form = ref({
    enter_number: "",
    enter_year: "",
    enter_subject: "",
    enter_source: "",
    enter_ratio: "",
    grade: "",
    post_name: "",
    post_code: "",
    workplace: "",
    post_category: "",
    perform_work: "",
    organization_name: "",
    organization_code: "",
    source_category: "",
    educational_require: "",
    degree_require: "",
    career: "",
    title_require: "",
    qualification: "",
    other: ""
})


const AuthorityUsers = ref([])
//获取部门用户
const departmentUinfo = async (Authority) => {
    const table = await getUserAuthorityList({ page: page.value, pageSize: pageSize.value, authority_id: Authority, ...searchInfo.value })
    if (table.code === 0) {
        console.log(table.data.list)

        AuthorityUsers.value = table.data.list
    }
}



const findPrincipal = (row, step) => {
    const setting = ref([])
    if (row.a3_setting !== '') {
        setting.value = JSON.parse(row.a3_setting)
        let name = setting.value.find(item => item.issue_number == step)
        return name.principal
    }
    return '暂无配置负责人'
}

const sourceCategory = ref([
    {
        value: '高校毕业生',
        label: '高校毕业生'
    },
    {
        value: '在校学生',
        label: '在校学生',
    },
    {
        value: '社会人士',
        label: '社会人士',
    }
])

const rules = ref({
    path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
    apiGroup: [
        { required: true, message: '请输入组名称', trigger: 'blur' }
    ],
    method: [
        { required: true, message: '请选择请求方式', trigger: 'blur' }
    ],
    description: [
        { required: true, message: '请输入api介绍', trigger: 'blur' }
    ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const onReset = () => {
    searchInfo.value = {}
}

const handleSuccess = (resp) => {
    if (resp.code === 0) {
        ElMessage({
            type: 'success',
            message: '图片上传成功',
            showClose: true
        })

        fileList.value.push({ name: resp.data.file.name, url: resp.data.file.url })
        form.value.photograph = JSON.stringify(fileList.value)
        return
    }
    ElMessage({
        type: 'error',
        message: '图片上传失败',
        showClose: true
    })

};

const handleRemove = (file, fileList) => {
    // 处理删除文件的逻辑，例如从文件列表中删除文件
    const index = fileList.indexOf(file);
    if (index !== -1) {
        fileList.splice(index, 1);
    }
    form.value.photograph = JSON.stringify(fileList.value)
};


// 搜索
const onSubmit = () => {
    page.value = 1
    pageSize.value = 10
    getTableData()
}

// 分页
const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
    if (prop) {
        if (prop === 'ID') {
            prop = 'id'
        }
        searchInfo.value.orderKey = toSQLLine(prop)
        searchInfo.value.desc = order === 'descending'
    }
    getTableData()
}

// 类别列表
const genreList = async () => {
    const table = await getGenreList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        genreList1.value = table.data.list
    }
}


// 类别列表
const supplier = async () => {
    const table = await getSupplierList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        supplierList.value = table.data.list
    }
}


// 类别列表
const project = async () => {
    const table = await getProjectList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        projectList.value = table.data.list
    }
}


// 查询
const getTableData = async () => {
    searchInfo.value.process_mode = "A3"
    const table = await getManageList({ page: page.value, pageSize: pageSize.value, orderKey: 'id', desc: true, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}


// 批量操作
const handleSelectionChange = (val) => {
    apis.value = val
}

// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
        enter_number: "",
        enter_year: "",
        enter_subject: "",
        enter_source: "",
        enter_ratio: "",
        grade: "",
        post_name: "",
        post_code: "",
        workplace: "",
        post_category: "",
        perform_work: "",
        organization_name: "",
        organization_code: "",
        source_category: "",
        educational_require: "",
        degree_require: "",
        career: "",
        title_require: "",
        qualification: "",
        other: ""
    }
}

const closeDialog = () => {
    initForm()
    drawer.value = false
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            const res = await updateManage(form.value)
            if (res.code === 0) {
                ElMessage({
                    type: 'success',
                    message: '岗位配置成功',
                    showClose: true
                })
            }
            getTableData()
            closeDialog()
        }
    })
}



</script>
<style lang="scss" scoped>
.drawer-container {
    transition: all 0.2s;

    &:hover {
        right: 0
    }

    & {
        position: fixed;
        right: -20px;
        bottom: 15%;
        height: 40px;
        width: 800px;
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 999;
        color: #fff;
        border-radius: 4px 0 0 4px;
        cursor: pointer;
        -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 10%);
    }

}

.setting_body {
    & {
        padding: 20px;
    }

    .setting_card {
        margin-bottom: 20px;
    }

    .setting_content {
        margin-top: 20px;
        display: flex;
        flex-direction: column;

        >.theme-box {
            display: flex;
        }

        >.color-box {
            div {
                display: flex;
                flex-direction: column;
            }
        }

        .item {
            & {
                display: flex;
                align-items: center;
                justify-content: center;
                flex-direction: column;
                margin-right: 20px;
            }


            .item-top {
                position: relative;
            }

            .check {
                position: absolute;
                font-size: 20px;
                color: #00afff;
                right: 10px;
                bottom: 10px;
            }

            p {
                text-align: center;
                font-size: 12px;
            }
        }
    }
}

.el-card__body {
    height: 800px;
}

.disabled {
    pointer-events: none;
    opacity: 0.5;
}
</style>