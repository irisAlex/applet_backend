<template>
    <div>
        <div class="gva-search-box">
            <a-form ref="searchForm" class="ant-advanced-search-form" :model="searchInfo" @finish="onFinish">
                <a-row :gutter="24">
                    <a-col :span="6">
                        <a-form-item label="项目名称">
                            <a-input v-model:value="searchInfo.name" placeholder="项目名称"></a-input>
                        </a-form-item>
                    </a-col>
                    <a-col :span="6">
                        <a-form-item label="项目周期">
                            <a-input v-model:value="searchInfo.period" placeholder="项目周期"></a-input>
                        </a-form-item>
                    </a-col>
                    <a-col :span="6">
                        <a-form-item label="负责人">
                            <a-input v-model:value="searchInfo.principal" placeholder="负责人"></a-input>
                        </a-form-item>
                    </a-col>
                    <a-col :span="6">
                        <a-form-item label="项目开始日期">
                            <a-date-picker v-model:value="searchInfo.created_at" placeholder="请选择日期" size="middle"
                                :locale="locale" />
                        </a-form-item>
                    </a-col>
                </a-row>
                <a-row>
                    <a-col :span="24" style="text-align: left;margin-bottom:20px;">
                        <a-button type="primary" html-type="submit" :icon="h(SearchOutlined)">
                            查询</a-button>
                        <a-button style="margin: 0 8px" @click="onReset" :icon="h(SyncOutlined)">重置</a-button>
                    </a-col>
                </a-row>
            </a-form>
        </div>

        <div class="gva-table-box">
            <div class="gva-btn-list">
                <a-button type="primary" @click="showModal('add')" :icon="h(PlusOutlined)">新增</a-button>
            </div>
            <a-table :columns="columns" :data-source="tableData" :scroll="{ x: 1500 }" :pagination="pagination"
                @change="handleTableChange">
                <template #bodyCell="{ column, record }">
                    <template v-if="column.key === 'operation'">
                        <a-button :icon="h(EditOutlined)" type="link" @click="editApiFunc(record)">修改</a-button>
                        <a-button :icon="h(DeleteOutlined)" type="link" @click="deleteApiFunc(record)">删除</a-button>
                    </template>
                </template>
            </a-table>
        </div>
        <div>
            <a-modal ref="modalRef" v-model:open="open" :wrap-style="{ overflow: 'hidden' }" @ok="handleOk"
                cancelText="取消" okText="确定" @cancel="Cancel">
                <a-form ref="formRef" :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
                    <a-form-item ref="name" label="项目名称">
                        <a-input v-model:value="formState.name" />
                    </a-form-item>
                    <a-form-item ref="principal" label="负责人">
                        <a-input v-model:value="formState.principal" />
                    </a-form-item>
                    <a-form-item ref="period" label="项目周期/天">
                        <a-row>
                            <a-col :span="12">
                                <a-slider v-model:value="formState.period" :min="1" :max="20" />
                            </a-col>
                            <a-col :span="4">
                                <a-input-number v-model:value="formState.period" :min="1" :max="100"
                                    style="margin-left: 16px" />
                            </a-col>
                        </a-row>
                    </a-form-item>
                    <a-form-item ref="priority" label="优先级">
                        <a-select v-model:value="formState.priority" style="width: 255px">
                            <a-select-option v-for="item in methodOptions" :key="item.value" :label="`${item.label}`"
                                :value="item.value" />
                        </a-select>
                    </a-form-item>
                    <a-form-item ref="describe" label="项目描述">
                        <a-textarea v-model:value="formState.describe" placeholder="项目描述" :rows="4" />
                    </a-form-item>
                </a-form>
                <template #title>
                    <div ref="modalTitleRef" style="width: 100%; cursor: move">{{ title }}</div>
                </template>
                <template #modalRender="{ originVNode }">
                    <div :style="transformStyle">
                        <component :is="originVNode" />
                    </div>
                </template>
            </a-modal>
        </div>
    </div>
</template>
<script lang='ts' setup>
import { ref, computed, watch, watchEffect, h, createVNode } from 'vue';
import { useDraggable } from '@vueuse/core';
import {
    SearchOutlined,
    SyncOutlined,
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    ExclamationCircleOutlined
} from '@ant-design/icons-vue';
import locale from 'ant-design-vue/es/date-picker/locale/zh_CN';
import {
    getProjectList,
    createProject,
    getProjectById,
    updateProject,
    deleteProject
} from '@/api/project'
import { message, Modal, TableColumnsType } from 'ant-design-vue';
const searchInfo = ref({
    name: '',
    principal: '',
    period: 0,
    priority: '',
    describe: ''
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

const labelCol = { span: 5 };
const wrapperCol = { span: 13 };
const columns: TableColumnsType = [
    { title: 'ID', width: 100, dataIndex: 'ID', key: 'name' },
    { title: '项目名称', width: 100, dataIndex: 'name', key: 'age' },
    { title: '描述', dataIndex: 'describe', key: 'addr', width: 150 },
    { title: '负责人', dataIndex: 'principal', key: 'addr', width: 150 },
    { title: '项目周期', dataIndex: 'period', key: 'contacts', width: 150 },
    { title: '优先级', dataIndex: 'priority', key: 'phone', width: 150 },
    {
        title: '操作',
        key: 'operation',
        fixed: 'right',
        width: 150,
    },
];
const searchForm = ref()
const title = ref('项目管理')
const open = ref(false);
const modalTitleRef = ref(null);
const type = ref()
const showModal = (e: any) => {
    switch (e) {
        case 'add':
            title.value = '项目管理'
            break
        case 'edit':
            title.value = '编辑项目'
            break
        default:
            break
    }
    type.value = e
    open.value = true;
};
const formRef = ref(null);
const formState = ref({
    name: '',
    principal: '',
    period: 0,
    priority: '',
    describe: ''
})
const editApiFunc = async (row: any) => {
    const res = await getProjectById({ id: row.ID })
    formState.value = res.data.project
    console.log(formState.value)
    showModal('edit')
}
const { x, y, isDragging } = useDraggable(modalTitleRef);

const Cancel = () => {
    formRef.value.resetFields()
    formState.value = {
        name: '',
        principal: '',
        period: 0,
        priority: '',
        describe: ''
    }
    open.value = false;
}
const startX = ref(0);
const startY = ref(0);
const startedDrag = ref(false);
const transformX = ref(0);
const transformY = ref(0);
const preTransformX = ref(0);
const preTransformY = ref(0);
const dragRect = ref({
    left: 0,
    right: 0,
    top: 0,
    bottom: 0,
});
watch([x, y], () => {
    if (!startedDrag.value) {
        startX.value = x.value;
        startY.value = y.value;
        const bodyRect = document.body.getBoundingClientRect();
        const titleRect = modalTitleRef.value.getBoundingClientRect();
        dragRect.value.right = bodyRect.width - titleRect.width;
        dragRect.value.bottom = bodyRect.height - titleRect.height;
        preTransformX.value = transformX.value;
        preTransformY.value = transformY.value;
    }
    startedDrag.value = true;
});
watch(isDragging, () => {
    if (!isDragging) {
        startedDrag.value = false;
    }
});
watchEffect(() => {
    if (startedDrag.value) {
        transformX.value =
            preTransformX.value +
            Math.min(Math.max(dragRect.value.left, x.value), dragRect.value.right) -
            startX.value;
        transformY.value =
            preTransformY.value +
            Math.min(Math.max(dragRect.value.top, y.value), dragRect.value.bottom) -
            startY.value;
    }
});
const transformStyle = computed(() => {
    return {
        transform: `translate(${transformX.value}px, ${transformY.value}px)`,
    };
});

const onFinish = (values: any) => {
    page.value = 1
    pageSize.value = 10
    getTableData()
};

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

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

const handleOk = async () => {
    switch (type.value) {
        case 'add':
            {
                const res = await createProject(formState.value)
                if (res.code === 0) {
                    message.success('添加成功');
                }
                getTableData()
                Cancel()
            }
            break
        case 'edit':
            {
                const res = await updateProject(formState.value)
                if (res.code === 0) {
                    message.success('数据更新成功');
                }
                getTableData()
                Cancel()
            }
            break
        default:
            {
                message.error('未知操作');
            }
            break
    }
}

const deleteApiFunc = async (row: any) => {
    Modal.confirm({
        title: '提示',
        icon: createVNode(ExclamationCircleOutlined),
        content: '此操作将永久删除数据, 是否继续?',
        okText: '确定',
        okType: 'danger',
        cancelText: '取消',
        onOk() {
            const deleteApi = async () => {
                const res = await deleteProject(row)
                if (res.code === 0) {
                    message.success('数据删除成功');
                    if (tableData.value.length === 1 && page.value > 1) {
                        page.value--
                    }
                    getTableData()
                }
            }
            deleteApi()
        },
        onCancel() {
            console.log('Cancel');
        },
    });
}

const onReset = () => {
    searchForm.value.resetFields()
    searchInfo.value = {
        name: '',
        principal: '',
        period: 0,
        priority: '',
        describe: ''
    }
}

const current = ref(1);
const pagination = computed(() => ({
    total: total.value,
    current: current.value,
    pageSize: pageSize.value,
    showSizeChanger: true,
    pageSizeOptions: ["10", "20", "50", "100"],
    showTotal: (total: any) => `共${total}条`,
}));

const handleTableChange = (pag: any) => {
    page.value = pag.current;
    current.value = pag.current;
    pageSize.value = pag.pageSize
    getTableData()
};

</script>

<style scoped>
#components-form-demo-advanced-search .ant-form {
    max-width: none;
}

#components-form-demo-advanced-search .search-result-list {
    margin-top: 16px;
    border: 1px dashed #e9e9e9;
    border-radius: 2px;
    background-color: #fafafa;
    min-height: 200px;
    text-align: center;
    padding-top: 80px;
}

[data-theme='dark'] .ant-advanced-search-form {
    background: rgba(255, 255, 255, 0.04);
    border: 1px solid #434343;
    padding: 24px;
    border-radius: 2px;
}

[data-theme='dark'] #components-form-demo-advanced-search .search-result-list {
    border: 1px dashed #434343;
    background: rgba(255, 255, 255, 0.04);
}

.code-box-demo .ant-slider {
    margin-bottom: 16px;
}
</style>