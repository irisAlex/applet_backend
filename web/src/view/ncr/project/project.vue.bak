<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="项目名称">
                    <el-input v-model="searchInfo.name" placeholder="项目名称" />
                </el-form-item>
                <el-form-item label="项目周期">
                    <el-input v-model.number="searchInfo.period" placeholder="项目周期" />
                </el-form-item>
                <el-form-item label="负责人">
                    <el-input v-model="searchInfo.principal" placeholder="项目名称" />
                </el-form-item>
                <el-form-item label="项目开始日期">
                    <el-date-picker v-model="searchInfo.created_at" type="date" placeholder="选择日期"
                        value-format="YYYY-MM-DDT15:04:05Z">
                    </el-date-picker>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog('addApi')">新增</el-button>
                <el-popover v-model="freshVisible" placement="top" width="160">
                    <p>确定要刷新Casbin缓存吗？</p>
                    <div style="text-align: right; margin-top: 8px;">
                        <el-button type="primary" link @click="freshVisible = false">取消</el-button>
                        <el-button type="primary" @click="onFresh">确定</el-button>
                    </div>
                    <template #reference>
                        <el-button icon="Refresh" @click="freshVisible = true">刷新缓存</el-button>
                    </template>
                </el-popover>
            </div>
            <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
                <el-table-column align="left" label="ID" min-width="150" prop="ID" sortable="custom" />
                <el-table-column align="left" label="项目名称" min-width="150" prop="name" sortable="custom" />
                <el-table-column align="left" label="描述" min-width="150" prop="describe" sortable="custom" />
                <el-table-column align="left" label="负责人" min-width="150" prop="principal" sortable="custom" />
                <el-table-column align="left" label="项目周期" min-width="150" prop="period" sortable="custom" />
                <el-table-column align="left" label="优先级" min-width="150" prop="priority" sortable="custom">
                </el-table-column>

                <el-table-column align="left" fixed="right" label="操作" width="300">
                    <template #default="scope">
                        <el-button icon="edit" type="primary" link @click="editApiFunc(scope.row)">修改</el-button>
                        <el-button icon="delete" type="primary" link @click="deleteApiFunc(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>

        </div>

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="40%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <el-form-item label="项目名称" prop="name" style="width:100%">
                    <el-input placeholder="项目名称" size="mini" v-model="form.name" />
                </el-form-item>
                <el-form-item label="负责人" prop="principal" style="width:100%">
                    <el-input placeholder="负责人" size="mini" v-model="form.principal" />
                </el-form-item>
                <el-form-item label="项目周期/天" prop="period" style="width:100%">
                    <el-slider v-model="form.period"></el-slider>
                </el-form-item>
                <el-form-item label="优先级" prop="priority" style="width:100%">
                    <el-select v-model="form.priority" placeholder="优先级" style="width:100%">
                        <el-option v-for="item in methodOptions" :key="item.value" :label="`${item.label}`"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="项目描述" prop="describe" style="width:100%">
                    <el-input type="textarea" placeholder="请输入内容" v-model="form.describe" maxlength="50" show-word-limit
                        :rows="10" />
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import {
    getProjectList,
    createProject,
    getProjectById,
    updateProject,
    deleteProject
} from '@/api/project'
import { toSQLLine } from '@/utils/stringFun'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { VideoCameraFilled } from '@element-plus/icons-vue'
import { toDoc } from '@/utils/doc'
import { deleteApisByIds } from '@/api/api';
// import moment from 'moment-timezone'

// //moment.tz.setDefault('Asia / Shanghai')
// moment.updateLocale('zh-cn', {
//     timezone: 'Asia/Shanghai'
// })

defineOptions({
    name: 'Api',
})

const methodFilter = (value) => {
    const target = methodOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
}

const apis = ref([])
const form = ref({
    period: 1,
    name: '',
    principal: '',
    describe: '',
    priority: ''
})
const methodOptions = ref([
    {
        value: 'P1',
        label: 'P1',
    },
    {
        value: 'P2',
        label: 'P2',
    },
    {
        value: 'P3',
        label: 'P3',
    },
    {
        value: 'P4',
        label: 'P4',
    },
    {
        value: 'P5',
        label: 'P5',
    },
    {
        value: 'P6',
        label: 'P6',
    },
    {
        value: 'P7',
        label: 'P7',
    },
    {
        value: 'P8',
        label: 'P8',
    },
    {
        value: 'P9',
        label: 'P9',
    }
])

const type = ref('')
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
const searchInfo = ref({
    //created_at: Date.now()
})

const onReset = () => {
    searchInfo.value = {}
}
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

// 查询
const getTableData = async () => {
    const table = await getProjectList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

// 批量操作
const handleSelectionChange = (val) => {
    apis.value = val
}

const deleteVisible = ref(false)
const onDelete = async () => {
    const ids = apis.value.map(item => item.ID)
    const res = await deleteApisByIds({ ids })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: res.msg
        })
        if (tableData.value.length === ids.length && page.value > 1) {
            page.value--
        }
        deleteVisible.value = false
        getTableData()
    }
}
const freshVisible = ref(false)
const onFresh = async () => {
    const res = await freshCasbin()
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: res.msg
        })
    }
    freshVisible.value = false
}

// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
        period: 1,
        name: '',
        principal: '',
        describe: '',
        priority: ''
    }
}

const pickerOptionsStart = (time) => {
    if (this.searchData.closeTime) {
        return time.getTime() >= new Date(this.searchData.closeTime).getTime()
    }
}

const dialogTitle = ref('项目管理')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
    switch (key) {
        case 'addApi':
            dialogTitle.value = '项目管理'
            break
        case 'edit':
            dialogTitle.value = '编辑Api'
            break
        default:
            break
    }
    type.value = key
    dialogFormVisible.value = true
}
const closeDialog = () => {
    initForm()
    dialogFormVisible.value = false
}

const editApiFunc = async (row) => {
    const res = await getProjectById({ id: row.ID })
    form.value = res.data.project
    openDialog('edit')
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            switch (type.value) {
                case 'addApi':
                    {
                        const res = await createProject(form.value)
                        if (res.code === 0) {
                            ElMessage({
                                type: 'success',
                                message: '添加成功',
                                showClose: true
                            })
                        }
                        getTableData()
                        closeDialog()
                    }

                    break
                case 'edit':
                    {
                        const res = await updateProject(form.value)
                        if (res.code === 0) {
                            ElMessage({
                                type: 'success',
                                message: '编辑成功',
                                showClose: true
                            })
                        }
                        getTableData()
                        closeDialog()
                    }
                    break
                default:
                    // eslint-disable-next-line no-lone-blocks
                    {
                        ElMessage({
                            type: 'error',
                            message: '未知操作',
                            showClose: true
                        })
                    }
                    break
            }
        }
    })
}

const deleteApiFunc = async (row) => {
    ElMessageBox.confirm('此操作将永久删除所有角色下该api, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    })
        .then(async () => {
            const res = await deleteProject(row)
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
</script>

<style scoped lang="scss">
.warning {
    color: #dc143c;
}
</style>