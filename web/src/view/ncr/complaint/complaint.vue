<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="反馈人">
                    <el-input v-model="searchInfo.interior_feedback_name" placeholder="反馈人" />
                </el-form-item>
                <el-form-item label="反馈单位">
                    <el-input v-model="searchInfo.interior_feedback_unit" placeholder="反馈单位" />
                </el-form-item>
                <el-form-item label="产品序列号">
                    <el-input v-model="searchInfo.product_sequence" placeholder="产品序列号" />
                </el-form-item>
                <el-form-item label="产品名称">
                    <el-input v-model="searchInfo.product_name" placeholder="产品名称" />
                </el-form-item>
                <el-form-item label="创建时间">
                    <el-date-picker v-model="value" type="date" placeholder="创建时间">
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
            </div>
            <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
                <el-table-column align="left" label="ID" min-width="150" prop="ID" sortable="custom" />
                <el-table-column align="left" label="产品名称" min-width="150" prop="product_name" sortable="custom" />
                <el-table-column align="left" label="产品序列号（物料编码）" min-width="230" prop="product_sequence"
                    sortable="custom" />
                <el-table-column align="left" label="客户名称" min-width="150" prop="client_name" sortable="custom" />
                <el-table-column align="left" label="投诉人" min-width="150" prop="complaint_name" sortable="custom" />
                <el-table-column align="left" label="项目" min-width="150" prop="project_name" sortable="custom" />
                <el-table-column align="left" label="状态" min-width="150" prop="status" sortable="custom" />
                <el-table-column align="left" label="受检物号" min-width="150" prop="checkout_number" sortable="custom" />
                <el-table-column align="left" label="内部反馈人" min-width="150" prop="interior_feedback_name"
                    sortable="custom" />
                <el-table-column align="left" label="内部反馈单位" min-width="150" prop="interior_feedback_unit"
                    sortable="custom" />
                <el-table-column align="left" label="问题描述（5W2H）附件照片" min-width="250" prop="issue_desc"
                    sortable="custom" />
                <el-table-column align="left" label="问题分级" min-width="150" prop="issue_level" sortable="custom" />
                <el-table-column align="left" label="短期措施/计划完成时间" min-width="250" prop="short_plan_date"
                    sortable="custom">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.short_plan_date) }}</span>
                    </template>
                </el-table-column>

                <el-table-column align="left" label="原因分析/附件/描述" min-width="230" prop="cause_desc" sortable="custom" />
                <el-table-column align="left" label="整改措施" min-width="150" prop="rectify" sortable="custom" />
                <el-table-column align="left" label="客服订单" min-width="150" prop="complaint_order" sortable="custom" />
                <el-table-column align="left" label="负责人" min-width="150" prop="principal_name" sortable="custom" />
                <el-table-column align="left" label="成本" min-width="150" prop="cost" sortable="custom" />
                <el-table-column align="left" label="整改计划完成时间" min-width="200" prop="rectify_date" sortable="custom">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.rectify_date) }}</span>
                    </template>
                </el-table-column>

                <el-table-column align="left" label="提交时间" min-width="150" prop="submit_date" sortable="custom">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.submit_date) }}</span>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="关闭时间" min-width="150" prop="close_date" sortable="custom">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.close_date) }}</span>
                    </template>
                </el-table-column>

                <el-table-column align="left" fixed="right" label="操作" width="300">
                    <template #default="scope">
                        <!-- <el-button icon="view" type="primary" link @click="editApiFunc(scope.row)">查看</el-button> -->
                        <el-button icon="edit" type="primary" link @click="editApiFunc(scope.row)">修改</el-button>
                        <el-button icon="delete" type="primary" link @click="deleteApiFunc(scope.row)">删除</el-button>
                        <el-button icon="circle-close" type="primary" link
                            @click="editApiFunc(scope.row)">关闭</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>

        </div>

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="70%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <el-form-item label="产品名称" prop="product_name" style="width:20%">
                    <el-input placeholder="产品名称" size="mini" v-model="form.product_name" />
                </el-form-item>
                <el-form-item label="产品序列号(物料编码)" prop="product_sequence" style="width:20%">
                    <el-input placeholder="产品序列号(物料编码)" size="mini" v-model="form.product_sequence" />
                </el-form-item>
                <el-form-item label="客户名称" prop="client_name" style="width:20%">
                    <el-input placeholder="客户名称" size="mini" v-model="form.client_name" />
                </el-form-item>
                <el-form-item label="投诉人" prop="complaint_name" style="width:20%">
                    <el-input placeholder="投诉人" size="mini" v-model="form.complaint_name" />
                </el-form-item>
                <el-form-item label="项目" prop="project_name" style="width:20%">
                    <el-input placeholder="项目" size="mini" v-model="form.project_name" />
                </el-form-item>
                <!-- <el-form-item label="状态" prop="status" style="width:20%">
                    <el-input placeholder="状态" size="mini" v-model="form.status" />
                </el-form-item> -->
                <el-form-item label="受检物号" prop="checkout_number" style="width:20%">
                    <el-input placeholder="受检物号" size="mini" v-model="form.checkout_number" />
                </el-form-item>
                <el-form-item label="内部反馈人" prop="interior_feedback_name" style="width:20%">
                    <el-input placeholder="内部反馈人" size="mini" v-model="form.interior_feedback_name" />
                </el-form-item>
                <el-form-item label="内部反馈单位" prop="interior_feedback_unit" style="width:20%">
                    <el-input placeholder="内部反馈单位" size="mini" v-model="form.interior_feedback_unit" />
                </el-form-item>
                <el-form-item label="负责人" prop="principal_name" style="width:20%">
                    <el-input placeholder="负责人" size="mini" v-model="form.principal_name" />
                </el-form-item>
                <el-form-item label="客服订单" prop="complaint_order" style="width:20%">
                    <el-input placeholder="客服订单" v-model="form.complaint_order" />
                </el-form-item>
                <el-form-item label="成本" prop="cost" style="width:20%">
                    <el-input placeholder="成本" v-model.number="form.cost" />
                </el-form-item>
                <el-form-item label="问题分级" prop="issue_level" style="width:20%">
                    <el-select v-model="form.issue_level" placeholder="问题分级" style="width:100%">
                        <el-option v-for="item in methodOptions" :key="item.value"
                            :label="`${item.label}(${item.value})`" :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="短期措施/计划完成时间" prop="short_plan_date" style="width:30%">
                    <el-date-picker v-model="form.short_plan_date" type="date" placeholder="计划完成时间"
                        value-format="YYYY-MM-DDT15:04:05Z">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="整改计划完成时间" prop="rectify_date" style="width:30%">
                    <el-date-picker v-model="form.rectify_date" type="date" placeholder="整改计划完成时间"
                        value-format="YYYY-MM-DDT15:04:05Z">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="提交时间" prop="submit_date" style="width:30%">
                    <el-date-picker v-model="form.submit_date" type="date" placeholder="整改计划完成时间"
                        value-format="YYYY-MM-DDT15:04:05Z">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="关闭时间" prop="close_date" style="width:30%">
                    <el-date-picker v-model="form.close_date" type="date" placeholder="关闭时间"
                        value-format="YYYY-MM-DDT15:04:05Z">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="问题描述（5W2H" prop="issue_desc" style="width:80%">
                    <el-input placeholder="问题描述（5W2H）" size="mini" type="textarea" v-model="form.issue_desc"
                        maxlength="50" show-word-limit :rows="10" />
                </el-form-item>
                <el-form-item label="问题描述附件照片" prop="issue_desc" style="width:80%">
                    <el-input placeholder="问题描述附件照片" size="mini" type="textarea" v-model="form.issue_desc"
                        maxlength="50" show-word-limit :rows="10" />
                </el-form-item>
                <el-form-item label="整改措施" prop="rectify" style="width:80%">
                    <el-input placeholder="整改措施" size="mini" type="textarea" v-model="form.rectify" maxlength="50"
                        show-word-limit :rows="10" />
                </el-form-item>
                <el-form-item label="原因分析描述" prop="cause_desc" style="width:100%">
                    <el-input placeholder="原因分析描述" size="mini" type="textarea" v-model="form.cause_desc" maxlength="50"
                        show-word-limit :rows="10" />
                </el-form-item>
                <el-form-item label="原因分析附件" prop="issue_desc" style="width:80%">
                    <el-input placeholder="原因分析附件" size="mini" type="textarea" v-model="form.issue_desc" maxlength="50"
                        show-word-limit :rows="10" />
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
    deleteComplaint,
    updateComplaint,
    getComplaintById,
    createComplaint,
    getComplaintlist,
    setPassDate
} from '@/api/complaint.js'
import { toSQLLine, formatDate } from '@/utils/stringFun'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
    name: 'Api',
})

const methodFilter = (value) => {
    const target = methodOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
}

const apis = ref([])
const form = ref({
    product_name: "",
    product_sequence: "",
    client_name: "",
    complaint_name: "",
    project_name: "",
    status: "",
    checkout_number: "",
    interior_feedback_name: "",
    interior_feedback_unit: "",
    issue_desc: "",
    issue_level: "",
    short_plan_date: "",
    cause_desc: "",
    complaint_order: "",
    principal_name: "",
    cost: "",
    rectify_date: "",
    submit_date: "",
    close_date: "",
    rectify: ""
})
const methodOptions = ref([
    {
        value: 'POST',
        label: '创建',
        type: 'success'
    },
    {
        value: 'GET',
        label: '查看',
        type: ''
    },
    {
        value: 'PUT',
        label: '更新',
        type: 'warning'
    },
    {
        value: 'DELETE',
        label: '删除',
        type: 'danger'
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
const searchInfo = ref({})

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
    const table = await getComplaintlist({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
        product_name: "",
        product_sequence: "",
        client_name: "",
        complaint_name: "",
        project_name: "",
        status: "",
        checkout_number: "",
        interior_feedback_name: "",
        interior_feedback_unit: "",
        issue_desc: "",
        issue_level: "",
        short_plan_date: "",
        cause_desc: "",
        complaint_order: "",
        principal_name: "",
        cost: "",
        rectify_date: "",
        submit_date: "",
        close_date: "",
        rectify: ""
    }
}

const dialogTitle = ref('添加客诉')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
    switch (key) {
        case 'addApi':
            dialogTitle.value = '添加客诉'
            break
        case 'edit':
            dialogTitle.value = '编辑客诉'
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
    const res = await getComplaintById({ id: row.ID })
    form.value = res.data.compliant
    openDialog('edit')
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            switch (type.value) {
                case 'addApi':
                    {
                        const res = await createComplaint(form.value)
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
                        const res = await updateComplaint(form.value)
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
            const res = await deleteComplaint(row)
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