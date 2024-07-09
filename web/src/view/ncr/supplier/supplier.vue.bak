<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="产品名称">
                    <el-input v-model="searchInfo.product" placeholder="产品名称" />
                </el-form-item>
                <el-form-item label="供应商地址">
                    <el-input v-model="searchInfo.addr" placeholder="供应商地址" />
                </el-form-item>
                <el-form-item label="供应商名称">
                    <el-input v-model="searchInfo.name" placeholder="供应商类型" />
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
                <el-table-column align="left" label="供应商名称" min-width="150" prop="name" sortable="custom" />
                <el-table-column align="left" label="供应商地址" min-width="150" prop="addr" sortable="custom" />
                <el-table-column align="left" label="供货产品名称" min-width="150" prop="product" sortable="custom" />
                <el-table-column align="left" label="联系人" min-width="150" prop="contacts" sortable="custom" />
                <el-table-column align="left" label="联系电话" min-width="150" prop="phone" sortable="custom" />
                <el-table-column align="left" label="Email" min-width="150" prop="email" sortable="custom" />
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

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="30%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <el-form-item label="供应商名称" prop="name" style="width:100%">
                    <el-input placeholder="供应商名称" autocomplete="off" v-model="form.name" />
                </el-form-item>
                <el-form-item label="供应商地址" prop="addr" style="width:100%">
                    <el-input placeholder="供应商地址" autocomplete="off" v-model="form.addr" />
                </el-form-item>
                <el-form-item label="供货产品名称" prop="product" style="width:100%">
                    <el-input placeholder="供货产品名称" autocomplete="off" v-model="form.product" />
                </el-form-item>
                <el-form-item label="联系人" prop="contacts" style="width:100%">
                    <el-input placeholder="联系人" autocomplete="off" v-model="form.contacts" />
                </el-form-item>
                <el-form-item label="联系电话" prop="phone" style="width:100%">
                    <el-input placeholder="联系电话" autocomplete="off" v-model="form.phone" />
                </el-form-item>
                <el-form-item label="Email" prop="email" style="width:100%">
                    <el-input placeholder="Email" autocomplete="off" v-model="form.email" />
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
    createSupplierApi,
    getSupplierList,
    getSupplierById,
    deleteSupplier,
    updateSupplier

} from '@/api/supplier'
import { toSQLLine } from '@/utils/stringFun'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
    name: 'Api',
})

const apis = ref([])
const form = ref({
    name: '',
    addr: '',
    email: '',
    phone: '',
    product: '',
    contacts: ''
})

const type = ref('')
const rules = ref({
    country: [{ required: true, message: '请输入国家', trigger: 'blur' }],
    genre: [
        { required: true, message: '请输入类别', trigger: 'blur' }
    ],
    name: [
        { required: true, message: '请输入名称', trigger: 'blur' }
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
    const table = await getSupplierList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
        name: '',
        addr: '',
        email: '',
        phone: '',
        product: '',
        contacts: ''
    }
}

const dialogTitle = ref('添加供应商')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
    switch (key) {
        case 'addApi':
            dialogTitle.value = '添加供应商'
            break
        case 'edit':
            dialogTitle.value = '编辑供应商'
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
    const res = await getSupplierById({ id: row.ID })
    form.value = res.data.supplier
    openDialog('edit')
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            switch (type.value) {
                case 'addApi':
                    {
                        const res = await createSupplierApi(form.value)
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
                        const res = await updateSupplier(form.value)
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
    ElMessageBox.confirm('此操作将永久删除数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    })
        .then(async () => {
            const res = await deleteSupplier(row)
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