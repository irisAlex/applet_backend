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
                    <el-input v-model="searchInfo.Post" placeholder="供应商名称" />
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
                <el-table-column align="left" label="招考人数" min-width="150" prop="serialnumber" />
                <el-table-column align="left" label="入围比例" min-width="150" prop="department" />
                <el-table-column align="left" label="招考年份" min-width="150" prop="mold" />
                <el-table-column align="left" label="进面分数线" min-width="150" prop="category" />
                <el-table-column align="left" label="数据来源" min-width="150" prop="checkout_name" />
                <el-table-column align="left" label="岗位名称" min-width="150" prop="checkout_number" />
                <el-table-column align="left" label="岗位代码" min-width="150" prop="process_mode" />
                <el-table-column align="left" label="岗位类别" min-width="150" prop="process_mode" />
                <el-table-column align="left" label="从事工作" min-width="150" prop="process_mode" />

                <el-table-column align="left" label="单位名称" min-width="150" prop="process_mode" />

                <el-table-column align="left" label="单位序号" min-width="150" prop="process_mode" />

                <el-table-column align="left" label="其他条件" min-width="150" prop="process_mode" />

                <el-table-column align="left" label="职业资格" min-width="150" prop="process_mode">
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
                <el-form ref="apiForm" :model="form" :inline="true">
                    <el-timeline>
                        <el-timeline-item timestamp="考情分析" placement="top">
                            <el-card>
                                <el-form-item label=" 招考人数:" prop="enter_number" style="width:50%">
                                    <el-input v-model="form.enter_number" @input='handleChange(1)'></el-input>
                                </el-form-item>

                                <el-form-item label="入围比例:" prop="enter_ratio" style="width:50%">
                                    <el-input v-model="form.enter_ratio" @input='handleChange(1)'></el-input>
                                </el-form-item>

                                <el-form-item label="招考年份:" prop="enter_year" style="width:50%">
                                    <el-date-picker v-model="form.enter_year" type="year" placeholder="选择年">
                                    </el-date-picker>
                                </el-form-item>

                                <el-form-item label="进面分数线:" prop="fractional_line" style="width:50%">
                                    <el-input v-model="form.fractional_line" @input='handleChange(1)'></el-input>
                                </el-form-item>

                                <el-form-item label="数据来源:" prop="enter_source" style="width:50%">
                                    <el-input v-model="form.enter_source" @input='handleChange(1)'></el-input>
                                </el-form-item>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="岗位信息" placement="top">
                            <el-card>
                                <el-form-item label="岗位名称:" prop="post_name" style="width:50%">
                                    <el-input v-model="form.post_name" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="岗位代码:" prop="post_code" style="width:50%">
                                    <el-input v-model="form.post_code" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="岗位类别:" prop="post_category" style="width:50%">
                                    <el-input v-model="form.post_category" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="从事工作:" prop="perform_work" style="width:50%">
                                    <el-input v-model="form.perform_work" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="单位名称:" prop="organization_name" style="width:50%">
                                    <el-input v-model="form.organization_name" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="单位序号:" prop="organization_code" style="width:50%">
                                    <el-input v-model="form.organization_code" @input='handleChange(1)'></el-input>
                                </el-form-item>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="报考条件" placement="top">
                            <el-card>
                                <el-form-item label="来源类别:" prop="source_category" style="width:80%">
                                    <el-select v-model="form.source_category" placeholder="请选择">
                                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="学历要求:" prop="educational_require" style="width:80%">
                                    <el-select v-model="form.educational_require" placeholder="请选择"
                                        @change="getSubjectNameList(form.educational_require)">
                                        <el-option v-for="item in education" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="专业科目:" prop="subject" style="width:50%">
                                    <el-select v-model="form.subject" placeholder="请选择"
                                        @change="getSubjectCareeNameList(form.subject)">
                                        <el-option v-for="item in careerList" :key="item.name" :label="item.name"
                                            :value="item.ID">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="所学专业:" prop="career" style="width:50%">
                                    <el-select v-model="form.career" placeholder="请选择"
                                        @change="getSpecialtyNameList(form.career)">
                                        <el-option v-for="item in specialtyList" :key="item.name" :label="item.name"
                                            :value="item.ID">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="专业名称:" prop="specialty" style="width:50%">
                                    <el-select v-model="form.specialty" placeholder="请选择">
                                        <el-option v-for="item in specialtyNameList" :key="item.name" :label="item.name"
                                            :value="item.name">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="学位要求:" prop="degree_require" style="width:50%">
                                    <el-input v-model="form.degree_require" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="职称要求:" prop="qualification" style="width:50%">
                                    <el-input v-model="form.qualification" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="职业资格:" prop="title_require" style="width:50%">
                                    <el-input v-model="form.title_require" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="其他条件:" prop="other" style="width:50%;">
                                    <el-input type="textarea" placeholder="请输入内容" v-model="form.other" maxlength="50"
                                        show-word-limit :rows="10" />
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
    getPostList,
    getSubjectByEId,
    getSubjectByEName
} from '@/api/post'
import { toSQLLine } from '@/utils/stringFun'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
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

const handleChange = (ste) => {
    setp.value = ste
}
const initDraw = () => {
}
const projectList = ref([])
const form = ref({
    enter_number: "",
    enter_year: "",
    enter_subject: "",
    enter_source: "",
    enter_ratio: "",
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
    other: "",
    specialty: "",
    subject: "",
    fractional_line
})


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

const education = ref([
    {
        value: '本科',
        label: '本科'
    },
    {
        value: '硕士',
        label: '硕士',
    },
    {
        value: '博士',
        label: '博士',
    },
    {
        value: '专科',
        label: '专科',
    }
])

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const onReset = () => {
    searchInfo.value = {}
}

const careerName = ref()
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

const specialtyList = ref([])
// 科目列表
const getSubjectCareeNameList = async (id) => {
    const table = await getSubjectByEId({ id: id })
    console.log(table, careerName)
    if (table.code === 0) {
        specialtyList.value = table.data.subject
    }
}

const specialtyNameList = ref([])
// 科目列表
const getSpecialtyNameList = async (id) => {
    const table = await getSubjectByEId({ id: id })
    console.log(table, careerName)
    if (table.code === 0) {
        specialtyNameList.value = table.data.subject
    }
}

const careerList = ref([])
// 科目列表
const getSubjectNameList = async (name) => {
    const table = await getSubjectByEName({ name: name })
    console.log(table)
    if (table.code === 0) {
        careerList.value = table.data.subject
    }
    console.log(careerList.value)
}
// 类别列表
const Post = async () => {
    const table = await getPostList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        PostList.value = table.data.post
    }
}

// getPostList()

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
        other: "",
        specialty: "",
        subject: "",
        fractional_line: ""
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