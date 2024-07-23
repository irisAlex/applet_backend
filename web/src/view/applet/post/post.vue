<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="部门" style="width:200px;" prop="method">
                </el-form-item>
                <el-form-item label="类型" style="width:200px;" prop="method">
                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                        <el-option v-for="item in sourceCategory" :key="item.value" :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="类别" prop="method" style="width:200px;">

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
                <el-table-column align="left" label="招考人数" min-width="150" prop="enter_number" />
                <el-table-column align="left" label="入围比例" min-width="150" prop="enter_ratio" />
                <el-table-column align="left" label="招考年份" min-width="150" prop="enter_year">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.enter_year) }}</span>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="进面分数线" min-width="150" prop="fractional_line" />
                <el-table-column align="left" label="数据来源" min-width="150" prop="enter_source" />
                <el-table-column align="left" label="岗位名称" min-width="150" prop="post_name" />
                <el-table-column align="left" label="岗位代码" min-width="150" prop="post_code" />
                <el-table-column align="left" label="岗位类别" min-width="150" prop="post_category" />
                <el-table-column align="left" label="从事工作" min-width="150" prop="perform_work" />

                <el-table-column align="left" label="单位名称" min-width="150" prop="organization_name" />

                <el-table-column align="left" label="单位序号" min-width="150" prop="organization_code" />

                <el-table-column align="left" label="其他条件" min-width="150" prop="other" />

                <el-table-column align="left" label="职业资格" min-width="150" prop="qualification">
                </el-table-column>
                <el-table-column align="left" fixed="right" label="操作" width="300">
                    <template #default="scope">
                        <el-button icon="edit" type="primary" link @click="modifyPostData(scope.row.ID)">修改</el-button>
                        <el-button type="primary" icon="delete" link @click="deleteApiFunc(scope.row)">
                            删除
                        </el-button>
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
                                    <el-input v-model.number="form.enter_number"></el-input>
                                </el-form-item>
                                <el-form-item label=" 报考人数:" prop="apply_number" style="width:50%">
                                    <el-input v-model.number="form.apply_number"></el-input>
                                </el-form-item>
                                <el-form-item label="入围比例:" prop="enter_ratio" style="width:50%">
                                    <el-input v-model="form.enter_ratio"></el-input>
                                </el-form-item>

                                <el-form-item label="招考年份:" prop="enter_year" style="width:50%">
                                    <el-date-picker v-model="form.enter_year" type="year" value-format="YYYY"
                                        placeholder="选择年">
                                    </el-date-picker>
                                </el-form-item>

                                <el-form-item label="进面分数线:" prop="fractional_line" style="width:50%">
                                    <el-input v-model="form.fractional_line"></el-input>
                                </el-form-item>

                                <el-form-item label="数据来源:" prop="enter_source" style="width:50%">
                                    <el-input v-model="form.enter_source"></el-input>
                                </el-form-item>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="岗位信息" placement="top">
                            <el-card>
                                <el-form-item label="岗位名称:" prop="post_name" style="width:50%">
                                    <el-input v-model="form.post_name"></el-input>
                                </el-form-item>
                                <el-form-item label="岗位代码:" prop="post_code" style="width:50%">
                                    <el-input v-model="form.post_code"></el-input>
                                </el-form-item>
                                <el-form-item label="岗位类别:" prop="post_category" style="width:50%">
                                    <el-input v-model="form.post_category"></el-input>
                                </el-form-item>
                                <el-form-item label="从事工作:" prop="perform_work" style="width:50%">
                                    <el-input v-model="form.perform_work"></el-input>
                                </el-form-item>
                                <el-form-item label="单位名称:" prop="organization_name" style="width:50%">
                                    <el-input v-model="form.organization_name"></el-input>
                                </el-form-item>
                                <el-form-item label="单位序号:" prop="organization_code" style="width:50%">
                                    <el-input v-model="form.organization_code"></el-input>
                                </el-form-item>
                                <el-form-item label="工作所在地:" prop="workplace" style="width:50%">
                                    <elui-china-area-dht v-model="form.workplace"></elui-china-area-dht>
                                </el-form-item>

                                <el-form-item label="户口类型:" prop="account" style="width:50%">
                                    <el-select v-model="form.account" placeholder="请选择">
                                        <el-option v-for="item in account" :key="item.value" :label="item.value"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
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
                                    <el-select v-model="form.degree_require" placeholder="请选择">
                                        <el-option v-for="item in degree" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="职称要求:" prop="qualification" style="width:50%">
                                    <el-input v-model="form.qualification"></el-input>
                                </el-form-item>
                                <el-form-item label="职业资格:" prop="title_require" style="width:50%">
                                    <el-input v-model="form.title_require"></el-input>
                                </el-form-item>
                                <el-form-item label="性别:" prop="gender" style="width:50%">
                                    <el-select v-model="form.gender" placeholder="请选择">
                                        <el-option v-for="item in gender" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="工作经验:" prop="work_experience" style="width:50%">
                                    <el-select v-model="form.work_experience" placeholder="请选择">
                                        <el-option v-for="item in work_experience" :key="item.value" :label="item.label"
                                            :value="item.value">
                                        </el-option>
                                    </el-select>
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
    getSubjectByEName,
    createPost,
    updatePost,
    getPostById,
    deletePost
} from '@/api/post'

import { toSQLLine, formatDate } from '@/utils/stringFun'
import { ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { EluiChinaAreaDht } from 'elui-china-area-dht'
defineOptions({
    name: 'Setting',
})

const drawer = ref(false)
const direction = ref('rtl')
const ize = ref('40%')
const dialogTitle = ref('岗位信息配置')
const typeT = ref("")
const openDialog = (key) => {
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
    typeT.value = key
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

const initDraw = () => {
}
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
    career_id: "",
    title_require: "",
    qualification: "",
    other: "",
    specialty: "",
    subject: "",
    subject_id: "",
    fractional_line: "",
    account: '',
    gender: '',
    work_experience: '',
    apply_number: '',
    province_name: '',
    province_code: '',
    city_code: '',
    district_code: ''
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

const gender = ref([
    {
        value: '男性',
    },
    {
        value: '女性',
    },
    {
        value: '性别不限',
    }
])

const work_experience = ref([
    {
        value: '要求工作经验',
    },
    {
        value: '工作经验不限',
    }
])


const education = ref([
    {
        value: '博士研究生',
        label: '博士研究生'
    },
    {
        value: '硕士研究生',
        label: '硕士研究生'
    },
    {
        value: '本科',
        label: '本科',
    },
    {
        value: '大专',
        label: '大专',
    },
    {
        value: '专科',
        label: '专科',
    },
    {
        value: '高中',
        label: '高中',
    },
    {
        value: '初中',
        label: '初中',
    },
    {
        value: '小学',
        label: '小学',
    }
])
const account = ref([
    {
        value: '农村',
        label: '农村'
    },
    {
        value: '城镇',
        label: '城镇',
    },
    {
        value: '外籍人士',
        label: '外籍人士',
    },
    {
        value: '户口不限',
        label: '户口不限',
    }
])

const degree = ref([
    {
        value: '博士',
        label: '博士'
    },
    {
        value: '硕士',
        label: '硕士',
    },
    {
        value: '学士',
        label: '学士',
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



watch(
    () => form.value.educational_require,
    () => {
        if (form.value.educational_require !== '') {
            getSubjectNameList(form.value.educational_require)

        }
    }
)

watch(
    () => form.value.subject_id,
    () => {
        if (form.value.subject_id !== '') {
            getSubjectCareeNameList(form.value.subject_id)
        }
    }
)


watch(
    () => form.value.career_id,
    () => {
        if (form.value.career_id !== '') {
            getSpecialtyNameList(form.value.career_id)
        }
    }
)

const modifyPostData = async (id) => {
    const table = await getPostById({ id: id })
    if (table.code !== 0) {
        {
            ElMessage({
                type: 'error',
                message: '网络消失在～～～～',
                showClose: true
            })
        }
        return
    }
    form.value = table.data.post
    form.value.workplace = JSON.parse(form.value.workplace)
    console.log(form.value.workplace)
    typeT.value = 'modify'
    drawer.value = true
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
    id = form.value.subject_id === '' ? id : form.value.subject_id
    if (id === '') {
        return
    }
    const table = await getSubjectByEId({ id: parseInt(id) })
    if (table.code !== 0) {
        {
            ElMessage({
                type: 'error',
                message: '网络消失在～～～～',
                showClose: true
            })
        }
        return
    }
    // form.value.specialty = ''
    // specialtyList.value = []
    specialtyList.value = table.data.subject
}

const specialtyNameList = ref([])
// 科目列表
const getSpecialtyNameList = async (id) => {
    id = form.value.career_id === '' ? id : form.value.career_id
    if (id === '') {
        return
    }
    const table = await getSubjectByEId({ id: parseInt(id) })
    if (table.code !== 0) {
        ElMessage({
            type: 'error',
            message: '网络消失在～～～～',
            showClose: true
        })
        return
    }
    // form.value.specialty = ''
    specialtyNameList.value = table.data.subject
}

const careerList = ref([])
// 科目列表
const getSubjectNameList = async (name) => {
    const table = await getSubjectByEName({ name: name })
    if (table.code !== 0) {
        ElMessage({
            type: 'error',
            message: '网络消失在～～～～',
            showClose: true
        })
        return
    }
    // form.value.subject = ''
    // form.value.specialty = ''
    // specialtyList.value = []
    // specialtyNameList.value = []
    careerList.value = table.data.subject
}
// 类别列表
const getTableData = async () => {
    const table = await getPostList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
    }
}

getTableData()

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
        career_id: "",
        title_require: "",
        qualification: "",
        other: "",
        specialty: "",
        subject: "",
        subject_id: "",
        fractional_line: "",
        account: '',
        gender: '',
        work_experience: '',
        apply_number: '',
        province_name: '',
        province_code: '',
        city_code: '',
        district_code: ''
    }
}

const closeDialog = () => {
    initForm()
    drawer.value = false
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            if (form.value.workplace.length >= 0) {
                province.value.forEach(item => {
                    if (item.code == form.value.workplace[0]) {
                        form.value.province_name = item.name
                    }
                })
                form.value.province_code = form.value.workplace[0]
                form.value.city_code = form.value.workplace[1]
                form.value.district_code = form.value.workplace[2]
                form.value.workplace = JSON.stringify(form.value.workplace)
            }
            switch (typeT.value) {
                case "setting":
                    {
                        form.value.career_id = form.value.career.toString()
                        form.value.subject_id = form.value.subject.toString()
                        form.value.subject = careerList.value.find(user => user.ID === form.value.subject).name
                        form.value.career = specialtyList.value.find(user => user.ID === form.value.career).name
                        const res = await createPost(form.value)
                        if (res.code === 0) {
                            ElMessage({
                                type: 'success',
                                message: '岗位配置成功',
                                showClose: true
                            })
                        }
                    }

                    break;
                case "modify":
                    {
                        const res = await updatePost(form.value)
                        if (res.code === 0) {
                            ElMessage({
                                type: 'success',
                                message: '岗位更新成功',
                                showClose: true
                            })
                        }
                    }
                    break
                default:
                    {
                        ElMessage({
                            type: 'error',
                            message: '未知操作',
                            showClose: true
                        })
                    }
                    break
            }

            getTableData()
            closeDialog()
        }
    })
}

const deleteApiFunc = async (row) => {
    ElMessageBox.confirm('此操作将永久删除该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    })
        .then(async () => {
            const res = await deletePost(row)
            if (res.code === 0) {
                ElMessage({
                    type: 'success',
                    message: '删除成功!'
                })
                if (tableData.value.length === 1 && page.value > 1) {
                    page.value--
                }
                getTableData()
            }
        })
}

const province = ref([
    {
        "code": "110000",
        "name": "北京市",
        "province": "11"
    },
    {
        "code": "120000",
        "name": "天津市",
        "province": "12"
    },
    {
        "code": "130000",
        "name": "河北省",
        "province": "13"
    },
    {
        "code": "140000",
        "name": "山西省",
        "province": "14"
    },
    {
        "code": "150000",
        "name": "内蒙古自治区",
        "province": "15"
    },
    {
        "code": "210000",
        "name": "辽宁省",
        "province": "21"
    },
    {
        "code": "220000",
        "name": "吉林省",
        "province": "22"
    },
    {
        "code": "230000",
        "name": "黑龙江省",
        "province": "23"
    },
    {
        "code": "310000",
        "name": "上海市",
        "province": "31"
    },
    {
        "code": "320000",
        "name": "江苏省",
        "province": "32"
    },
    {
        "code": "330000",
        "name": "浙江省",
        "province": "33"
    },
    {
        "code": "340000",
        "name": "安徽省",
        "province": "34"
    },
    {
        "code": "350000",
        "name": "福建省",
        "province": "35"
    },
    {
        "code": "360000",
        "name": "江西省",
        "province": "36"
    },
    {
        "code": "370000",
        "name": "山东省",
        "province": "37"
    },
    {
        "code": "410000",
        "name": "河南省",
        "province": "41"
    },
    {
        "code": "420000",
        "name": "湖北省",
        "province": "42"
    },
    {
        "code": "430000",
        "name": "湖南省",
        "province": "43"
    },
    {
        "code": "440000",
        "name": "广东省",
        "province": "44"
    },
    {
        "code": "450000",
        "name": "广西壮族自治区",
        "province": "45"
    },
    {
        "code": "460000",
        "name": "海南省",
        "province": "46"
    },
    {
        "code": "500000",
        "name": "重庆市",
        "province": "50"
    },
    {
        "code": "510000",
        "name": "四川省",
        "province": "51"
    },
    {
        "code": "520000",
        "name": "贵州省",
        "province": "52"
    },
    {
        "code": "530000",
        "name": "云南省",
        "province": "53"
    },
    {
        "code": "540000",
        "name": "西藏自治区",
        "province": "54"
    },
    {
        "code": "610000",
        "name": "陕西省",
        "province": "61"
    },
    {
        "code": "620000",
        "name": "甘肃省",
        "province": "62"
    },
    {
        "code": "630000",
        "name": "青海省",
        "province": "63"
    },
    {
        "code": "640000",
        "name": "宁夏回族自治区",
        "province": "64"
    },
    {
        "code": "650000",
        "name": "新疆维吾尔自治区",
        "province": "65"
    },
    {
        "code": "710000",
        "name": "台湾省",
        "province": "71"
    },
    {
        "code": "810000",
        "name": "香港特别行政区",
        "province": "81"
    },
    {
        "code": "820000",
        "name": "澳门特别行政区",
        "province": "82"
    }
])

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