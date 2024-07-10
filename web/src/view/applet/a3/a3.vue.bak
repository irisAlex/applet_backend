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
                        <el-option v-for="item in moldList" :key="item.value" :label="item.label" :value="item.value">
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
                    <el-timeline v-if="isShowReport">
                        <el-timeline-item timestamp="1.问题简述&组建团队" placement="top" :class="isDisabled1" v-if="isShow1">
                            <el-card>
                                <el-form-item label=" 区域/位置:" prop="area" style="width:20%">
                                    <span>{{ form.area }}</span>
                                </el-form-item>

                                <el-form-item label="工序/工位:" prop="station" style="width:20%">
                                    <span>{{ form.station }}</span>
                                </el-form-item>

                                <el-form-item label="零件名称:" prop="checkout_name" style="width:20%">
                                    <span>{{ form.checkout_name }}</span>
                                </el-form-item>

                                <el-form-item label="零件编号:" prop="serialnumber" style="width:20%">
                                    <span>{{ form.serialnumber }}</span>
                                </el-form-item>

                                <el-form-item label="缺陷描述:" prop="defect_problem" style="width:20%">
                                    <span>{{ form.defect_problem }}</span>
                                </el-form-item>

                                <el-form-item label="零件数量:" prop="packages_number" style="width:20%">
                                    <span>{{ form.packages_number }}</span>
                                </el-form-item>

                                <el-form-item label="发现地点:" prop="find_addr" style="width:20%">
                                    <span>{{ form.find_addr }}</span>
                                </el-form-item>

                                <el-form-item label="发现工序:" prop="find_process" style="width:20%">
                                    <span>{{ form.find_process }}</span>
                                </el-form-item>

                                <el-form-item label="问题描述:" prop="describe" style="width:100%">
                                    <div>
                                        <QuillEditor :options="editorOptions" content-type="html" ref="quillEditor"
                                            theme="snow" v-model:content="form.describe" :value="form.describe"
                                            readOnly />
                                    </div>
                                </el-form-item>
                                <el-form-item label="组长:" prop="describe" style="width:40%">
                                    <el-input v-model="form.a3_team_top" @input='handleChange(1)'></el-input>
                                </el-form-item>
                                <el-form-item label="团队成员:" prop="describe" style="width:40%">
                                    <el-input v-model="form.a3_team_member" @input='handleChange(1)'></el-input>
                                </el-form-item>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="2.问题描述" placement="top" :class="isDisabled2" v-if="isShow2">
                            <el-card>
                                <el-table :data="tableIssueData" border style="width: 100%;" :stripe="true"
                                    :disabled="true">
                                    <el-table-column label="描述" align="center" min-width="150"
                                        prop="product_serialnumber">
                                        <template #default="scope">
                                            <span>{{ scope.row.describe }}</span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="问题" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.issue" @input='handleChange(2)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="(在其他产品/过程)是否有问题" min-width="150" prop="w_serialnumber">
                                        <template #default="scope">
                                            <el-switch v-model="scope.row.isIssue" active-color="#13ce66"
                                                inactive-color="#D3D3D3">
                                            </el-switch>
                                        </template>
                                    </el-table-column>
                                </el-table>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="3.应急响应措施" placement="top" :class="isDisabled3" v-if="isShow3">
                            <el-card>
                                <el-table :data="tableEmergencyData" border style="width: 100%;" :stripe="true">
                                    <el-table-column label="NO" align="center" min-width="150"
                                        prop="product_serialnumber">
                                        <template #default="scope">
                                            <span> {{ scope.$index + 1 }}</span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="应急相应措施(围堵)" align="center" min-width="150" prop="scheme">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.scheme" @input='handleChange(3)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="负责人" align="center" min-width="150" prop="name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.name"></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="日期" align="center" min-width="150" prop="date">
                                        <template #default="scope">
                                            <el-date-picker v-model="scope.row.date" type="date" placeholder="选择日期"
                                                value-format="YYYY-MM-DDT15:04:05Z">
                                            </el-date-picker>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="状态" align="center" min-width="150" prop="state">
                                        <template #default="scope">
                                            <Cri :defaultValue="scope.row.state" v-model="scope.row.state" />
                                        </template>
                                    </el-table-column>
                                </el-table>
                                <el-icon @click="addTableData('emergency')" class="icon" size="24" color="#fb7a14">
                                    <CirclePlusFilled />
                                </el-icon>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="4.原因分析" placement="top" :class="isDisabled4" v-if="isShow4">
                            <el-card>
                                <FishBone :height="500" @click-node="onClickNode" :id="ID" @issues="onIssues"
                                    @reduceIssue="onReduceIssue" :step="setp" />
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="5.根本原因确认" placement="top" :class="isDisabled5" v-if="isShow5">
                            <el-card>
                                <el-table :data="tableCauseData" border style="width: 100%;" :stripe="true">
                                    <el-table-column label="" align="center" min-width="50" prop="product_serialnumber">
                                        <template #default="scope">
                                            <span>why?</span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column v-for="(issue, index) in Issues" :key="index"
                                        :label="(index + 1) + '.' + issue.name + '?'" align="center" min-width="150"
                                        prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row[scope.$index + index]"
                                                @input='handleChange(5)'></el-input>
                                        </template>
                                    </el-table-column>
                                </el-table>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="6.纠正措施" placement="top" :class="isDisabled6" v-if="isShow6">
                            <el-card>
                                <el-table :data="tableCorrectData" border style="width: 100%;" :stripe="true">
                                    <el-table-column label="NO" align="center" min-width="150"
                                        prop="product_serialnumber">
                                        <template #default="scope">
                                            <span> {{ scope.$index + 1 }}</span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="确定纠正措施" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.scheme" @input='handleChange(6)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="牵头人" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.name" @input='handleChange(6)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="日期" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-date-picker v-model="scope.row.date" type="date" placeholder="选择日期"
                                                value-format="YYYY-MM-DDT15:04:05Z">
                                            </el-date-picker>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="状态" min-width="150" prop="w_serialnumber">
                                        <template #default="scope">
                                            <Cri :defaultValue="scope.row.state" v-model="scope.row.state" />
                                        </template>
                                    </el-table-column>
                                </el-table>
                                <el-icon @click="addTableData('correct')" class="icon" size="24" color="#fb7a14">
                                    <CirclePlusFilled />
                                </el-icon>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="7.措施验证" placement="top" :class="isDisabled7" v-if="isShow7">
                            <el-card>
                                <el-table :data="tableDeliveryData" border style="width: 100%;" :stripe="true">
                                    <el-table-column label="NO" align="center" min-width="150"
                                        prop="product_serialnumber">
                                        <template #default="scope">
                                            <span> {{ scope.$index + 1 }}</span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="检验/验证" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.scheme" @input='handleChange(7)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="证据" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-upload action="/api/fileUploadAndDownload/upload"
                                                list-type="picture-card" :on-preview="handlePictureCardPreview"
                                                :on-remove="handleRemove" :headers="{ 'x-token': userStore.token }">
                                                <i class="el-icon-plus"></i>
                                            </el-upload :headers="{ 'x-token': userStore.token }">
                                            <el-dialog v-model="dialogFormVisible">
                                                <img width="100%" :src="dialogImageUrl" />
                                            </el-dialog>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="牵头人" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.name" @input='handleChange(7)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="日期" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-date-picker v-model="scope.row.date" type="date" placeholder="选择日期"
                                                value-format="YYYY-MM-DDT15:04:05Z">
                                            </el-date-picker>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="状态" min-width="150" prop="w_serialnumber">
                                        <template #default="scope">
                                            <Cri :defaultValue="scope.row.state" v-model="scope.row.state" />
                                        </template>
                                    </el-table-column>
                                </el-table>
                                <el-icon @click="addTableData('delivery')" class="icon" size="24" color="#fb7a14">
                                    <CirclePlusFilled />
                                </el-icon>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="8.保持改进结果" placement="top" :class="isDisabled8" v-if="isShow8">
                            <el-card>
                                <el-table :data="tableResultData" border style="width: 100%;" :stripe="true">
                                    <el-table-column label="NO" align="center" min-width="150"
                                        prop="product_serialnumber">
                                        <template #default="scope">
                                            <span> {{ scope.$index + 1 }}</span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="保持改进结果" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.scheme" @input='handleChange(8)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="牵头人" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.name" @input='handleChange(8)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="日期" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-date-picker v-model="scope.row.date" type="date" placeholder="选择日期"
                                                value-format="YYYY-MM-DDT15:04:05Z">
                                            </el-date-picker>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="状态" min-width="150" prop="w_serialnumber">
                                        <template #default="scope">
                                            <Cri :defaultValue="scope.row.state" v-model="scope.row.state" />
                                        </template>
                                    </el-table-column>
                                </el-table>
                                <el-icon @click="addTableData('result')" class="icon" size="24" color="#fb7a14">
                                    <CirclePlusFilled />
                                </el-icon>
                            </el-card>
                        </el-timeline-item>
                        <el-timeline-item timestamp="9.传递" placement="top" :class="isDisabled9" v-if="isShow9">
                            <el-card>
                                <el-table :data="tableEndData" border style="width: 100%;" :stripe="true">
                                    <el-table-column label="NO" align="center" min-width="150"
                                        prop="product_serialnumber">
                                        <template #default="scope">
                                            <span> {{ scope.$index + 1 }}</span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="标的/活动" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.scheme" @input='handleChange(9)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="牵头人" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.name" @input='handleChange(9)'></el-input>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="日期" align="center" min-width="150" prop="product_name">
                                        <template #default="scope">
                                            <el-date-picker v-model="scope.row.date" type="date" placeholder="选择日期"
                                                value-format="YYYY-MM-DDT15:04:05Z">
                                            </el-date-picker>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="状态" min-width="150" prop="w_serialnumber">
                                        <template #default="scope">
                                            <Cri :defaultValue="scope.row.state" v-model="scope.row.state" />
                                        </template>
                                    </el-table-column>
                                </el-table>
                                <el-icon @click="addTableData('end')" class="icon" size="24" color="#fb7a14">
                                    <CirclePlusFilled />
                                </el-icon>
                            </el-card>
                        </el-timeline-item>
                    </el-timeline>
                    <el-timeline v-if="!isShowReport">
                        <el-timeline-item v-for="(active, index) in configure" :key="index" :timestamp="active.title"
                            placement="top">
                            <el-card>
                                <el-form :inline="true">
                                    <el-form-item label="部门:" prop="serialnumber" style="width:25%">
                                        <el-select v-model="active.department" placeholder="选择部门"
                                            @change="departmentUinfo(active.department)">
                                            <el-option v-for="item in departmentList" :key="item.authorityId"
                                                :label="item.authorityName" :value="item.authorityId">
                                            </el-option>
                                        </el-select>
                                    </el-form-item>
                                    <el-form-item label="负责人:" prop="notification_member" style="width:20%">
                                        <el-select v-model="active.principal" placeholder="选择人员" style="width:100%">
                                            <el-option v-for="item in AuthorityUsers" :key="item.ID"
                                                :label="`${item.userName}`" :value="item.userName" />
                                        </el-select>
                                    </el-form-item>
                                    <el-form-item label="限定时间/天:" prop="limit_date" style="width:30%">
                                        <el-input placeholder="限定时间" size="mini" v-model="active.limit_date" />
                                    </el-form-item>
                                </el-form>
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
    getAuthorityList,
    getGenreList,
    getSupplierList,
    getProjectList,
    updateManage,
    getManageById,
    getManageList,
    downFile,
    getUserAuthorityList,
    down
} from '@/api/manage.js'
import { toSQLLine, formatDate } from '@/utils/stringFun'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { strToJson } from '@/utils/format'
import FishBone from '@/components/fishBone/index.vue'
import { useUserStore } from '@/pinia/modules/user'
import Cri from '@/components/circle/index.vue'
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
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

const cir = ref(0)
const stateAdd = () => {
    return cir.value = 1.5
}

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
const ize = ref('70%')
const isShowReport = ref(false)
const dialogTitle = ref('A3配置')
const openDialog = (row, key) => {
    switch (key) {
        case 'setting':
            dialogTitle.value = 'A3配置'
            isShowReport.value = false
            break
        case 'report':
            dialogTitle.value = 'A3填写报告'
            isShowReport.value = true
            break
        default:
            break
    }
    showSettingDrawer(key, row)
}

function onClickNode(data, imgbase4) {
    form.value.a3_img_base64 = imgbase4
    form.value.a3_affirm = JSON.stringify(data)
    isPass.value = 4
}

const Issues = ref([])
function onIssues(data) {
    Issues.value.push(data)
}

function onReduceIssue(data) {
    const target = Issues.value.findIndex(item => {
        return item.name === data
    })
    Issues.value.splice(target, 1)
}


const handleClose = () => {
    initForm()
    initDraw()
    drawer.value = false
}
const ID = ref(1)
const showSettingDrawer = async (key, row) => {
    const res = await getManageById({ id: row.ID })
    ID.value = row.ID
    if (res.code === 0) {
        form.value = res.data.manage
        tableIssueData.value = strToJson(form.value.a3_issue_desc, tableIssueData.value)
        tableCauseData.value = strToJson(form.value.a3_cause, tableCauseData.value)
        tableResultData.value = strToJson(form.value.a3_result, tableResultData.value)
        tableCorrectData.value = strToJson(form.value.a3_correct, tableCorrectData.value)
        tableEmergencyData.value = strToJson(form.value.a3_contingency, tableEmergencyData.value)
        tableEndData.value = strToJson(form.value.a3_end, tableEndData.value)
        tableDeliveryData.value = strToJson(form.value.a3_verify, tableDeliveryData.value)
        configure.value = strToJson(form.value.a3_setting, configure.value)
        Issues.value = form.value.a3_issues === '' ? [] : JSON.parse(form.value.a3_issues)
        setp.value = form.value.a3_step
    }
    if (form.value.a3_setting == '' && key !== 'setting') {
        ElMessage({
            type: 'error',
            message: '请先配置填写报告人员',
            showClose: true
        })
        return
    }
    permissionProcess()
    drawer.value = true
}
const setp = ref(0)
const isPass = ref(0)
const imgBase64 = ref('')
const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'

const downExcel = async (row) => {
    // const res = await down()
    // let blob = new Blob([res.data], {
    //     //因为我实例传输的是csv格式的，按照需求改动
    //     type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
    // });
    // let a = document.createElement('a')
    // // 兼容IE
    // if (!!window.ActiveXObject || "ActiveXObject" in window) {
    //     // IE
    //     window.navigator.msSaveBlob(blob, "字典列表.txt");
    // } else {
    //     // 非IE
    //     a.setAttribute("download", "A3.xlsx");
    // }
    // //这边就是模拟个a标签，让他来点击，因为原版的a标签虽然可以实现，但是是为get方法，且不能携带token等信息
    // a.href = window.URL.createObjectURL(blob)
    // document.body.appendChild(a)
    // a.click()
    // URL.revokeObjectURL(a.href)
    // document.body.removeChild(a)


    const res = await downFile({ id: row.ID })
    if (res.code === 0) {
        if (res.data.manage.a3_step < 9) {
            ElMessage({
                type: 'error',
                message: '请先填写完整报告，在下载。',
                showClose: true
            })
            return
        }
        window.open(path + "uploads/file/" + row.serialnumber + "_A3.xlsx")
    }
}

const isDisabled9 = ref('disabled')
const isDisabled1 = ref('disabled')
const isDisabled2 = ref('disabled')
const isDisabled3 = ref('disabled')
const isDisabled4 = ref('disabled')
const isDisabled5 = ref('disabled')
const isDisabled6 = ref('disabled')
const isDisabled7 = ref('disabled')
const isDisabled8 = ref('disabled')

const isShow9 = ref(false)
const isShow1 = ref(false)
const isShow2 = ref(false)
const isShow3 = ref(false)
const isShow4 = ref(false)
const isShow5 = ref(false)
const isShow6 = ref(false)
const isShow7 = ref(false)
const isShow8 = ref(false)

const userStore = useUserStore()

const permissionProcess = () => {
    for (let item of configure.value) {
        if (item.issue_number <= (setp.value + 1)) {
            const tmp = 'isShow' + item.issue_number.toString()
            const tmp1 = 'isDisabled' + item.issue_number.toString()
            eval(tmp).value = true
            if (userStore.userInfo.userName === item.principal) {
                eval(tmp1).value = ''
            }
        }
    }
}


const handleChange = (ste) => {
    setp.value = ste
}
const initDraw = () => {
    setp.value = 0
    isDisabled9.value = 'disabled'
    isDisabled1.value = 'disabled'
    isDisabled2.value = 'disabled'
    isDisabled3.value = 'disabled'
    isDisabled4.value = 'disabled'
    isDisabled5.value = 'disabled'
    isDisabled6.value = 'disabled'
    isDisabled7.value = 'disabled'
    isDisabled8.value = 'disabled'

    isShow9.value = false
    isShow1.value = false
    isShow2.value = false
    isShow3.value = false
    isShow4.value = false
    isShow5.value = false
    isShow6.value = false
    isShow7.value = false
    isShow8.value = false

    tableCorrectData.value = ([
        {
            name: '',
            date: '',
            scheme: "",
            state: "1"
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: "1"
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: "1"
        }
    ])



    tableCauseData.value = ([
        {},
        {},
        {},
        {},
        {}
    ])
    tableIssueData.value = ([{
        describe: "(what)这是什么问题？",
        issue: "",
        isIssue: false
    }, {
        describe: "(where)问题在哪里出现？",
        issue: "",
        isIssue: false
    },
    {
        describe: "(when)问题发生的时间？",
        issue: "",
        isIssue: false
    },
    {
        describe: "(why)为什么说这是一个问题？",
        issue: "",
        isIssue: false
    },
    {
        describe: "(who)谁发现&谁制造？",
        issue: "",
        isIssue: false
    },
    {
        describe: "(how)问题如何被发现？",
        issue: "",
        isIssue: false
    },
    {
        describe: "(how much/many)程度/数量",
        issue: "",
        isIssue: false
    },
    ])


    //问题3
    tableEmergencyData.value = ([
        {
            name: '',
            date: '',
            scheme: "",
            state: 1
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: 1
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: 1
        }
    ])

    //问题6
    tableDeliveryData.value = ([
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        }
    ])
    //问题7
    tableResultData.value = ([
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        }
    ])

    //问题7
    tableEndData.value = ([
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        },
        {
            name: '',
            date: '',
            scheme: "",
            state: ""
        }
    ])

}
const tableIssueData = ref([{
    describe: "(what)这是什么问题？",
    issue: "",
    isIssue: false
}, {
    describe: "(where)问题在哪里出现？",
    issue: "",
    isIssue: false
},
{
    describe: "(when)问题发生的时间？",
    issue: "",
    isIssue: false
},
{
    describe: "(why)为什么说这是一个问题？",
    issue: "",
    isIssue: false
},
{
    describe: "(who)谁发现&谁制造？",
    issue: "",
    isIssue: false
},
{
    describe: "(how)问题如何被发现？",
    issue: "",
    isIssue: false
},
{
    describe: "(how much/many)程度/数量",
    issue: "",
    isIssue: false
},
])




//问题3
const tableEmergencyData = ref([
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    }
])

//问题6
const tableDeliveryData = ref([
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    }
])
//问题7
const tableResultData = ref([
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    }
])

//问题7
const tableEndData = ref([
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    }
])

const addTableData = (row) => {
    const newRow = {
        state: 1
    }

    switch (row) {
        case 'emergency':
            tableEmergencyData.value.push(newRow)
            break
        case 'correct':
            tableCorrectData.value.push(newRow)
            break
        case 'delivery':
            tableDeliveryData.value.push(newRow)
            break
        default:
            break
    }
}
//问题4
const tableCorrectData = ref([
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    },
    {
        name: '',
        date: '',
        scheme: "",
        state: ""
    }
])



const tableCauseData = ref([
    {

    },
    {

    },
    {

    },
    {

    },
    {

    }
])



const configure = ref([
    {
        title: "1.发现问题",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 1
    },
    {
        title: "2.问题描述",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 2
    },
    {
        title: "3.应急响应措施",
        principal: "",
        limit_date: 0,
        department: "",
        issue_number: 3
    },
    {
        title: "4.原因分析",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 4
    },
    {
        title: "5.根本原因确认",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 5
    },
    {
        title: "6.纠正措施",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 6
    },
    {
        title: "7.措施验证",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 7
    },
    {
        title: "8.保持改进结果",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 8
    }, {
        title: "9.传递",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 9
    }
])

const apis = ref([])
const departmentList = ref([])
const genreList1 = ref([])
const supplierList = ref([])
const projectList = ref([])
const fileList = ref([])
const form = ref({
    serialnumber: "",
    department: "",
    mold: "",
    category: "",
    project: "",
    checkout_name: "",
    checkout_number: "",
    graph_number: "",
    version_number: "",
    purchase_order: "",
    production_order: "",
    delivery_order: "",
    packages_number: "",
    reject_packages_number: "",
    sample_checkout_number: "",
    reject_sample_checkout_number: "",
    supplier: "",
    checkout_date: "",
    describe: "",
    photograph: "",
    process_mode: "",
    duty_department: "",
    cause_desc: "",
    fill_from_date: "",
    disposal_concept: "",
    rework_number: 0,
    rework_man_hour: 0,
    rework_quantities: "",
    rework_process: "",
    rework_plan_date: "0001-01-01T00:00:00Z",
    rework_desc: "",
    rework_attachment: "",
    repair_plan_date: "0001-01-01T00:00:00Z",
    repair_desc: "",
    repair_attachment: "",
    parts_desc: "",
    series: "",
    operation_type: "",
    pass_date: "0001-01-01T00:00:00Z",
    area: "",
    find_addr: "",
    find_process: "",
    defect_problem: "",
    station: "",
    a3_principal: "",
    a3_team_top: "",
    a3_team_member: "",
    a3_issue_desc: "",
    a3_contingency: "",
    a3_cause: "",
    a3_affirm: "",
    a3_correct: "",
    a3_verify: "",
    a3_result: "",
    a3_end: "",
    a3_setting: "",
    a3_step: 0,
    a3_issues: '',
    a3_img_base64: ''
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

const methodOptions = ref([
    {
        value: 'Just Do it',
        label: 'Just Do it',
        type: 'success'
    },
    {
        value: 'A3',
        label: 'A3',
        type: ''
    },
    {
        value: '8D',
        label: '8D',
        type: 'warning'
    }
])

const moldList = ref([
    {
        value: '内部',
        label: '内部',
        type: 'success'
    },
    {
        value: '外部',
        label: '外部',
        type: ''
    },
    {
        value: '配做',
        label: '配做',
        type: 'warning'
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

// 部门列表
const department = async () => {
    const table = await getAuthorityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        departmentList.value = table.data.list
    }
}

department()

// 类别列表
const genreList = async () => {
    const table = await getGenreList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        genreList1.value = table.data.list
    }
}

genreList()

// 类别列表
const supplier = async () => {
    const table = await getSupplierList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        supplierList.value = table.data.list
    }
}

supplier()

// 类别列表
const project = async () => {
    const table = await getProjectList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        projectList.value = table.data.list
    }
}

project()

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
        serialnumber: "",
        department: "",
        mold: "",
        category: "",
        project: "",
        checkout_name: "",
        checkout_number: "",
        graph_number: "",
        version_number: "",
        purchase_order: "",
        production_order: "",
        delivery_order: "",
        packages_number: "",
        reject_packages_number: "",
        sample_checkout_number: "",
        reject_sample_checkout_number: "",
        supplier: "",
        checkout_date: "",
        describe: "",
        photograph: "",
        process_mode: "",//处事方式
        duty_department: "",
        cause_desc: "", //原因分析
        fill_from_date: "", //填表日期
        disposal_concept: "",
        rework_number: 0,
        rework_man_hour: 0,
        rework_quantities: "",
        rework_process: "",
        rework_plan_date: "0001-01-01T00:00:00Z",
        rework_desc: "",
        rework_attachment: "",
        repair_plan_date: "0001-01-01T00:00:00Z",
        repair_desc: "",
        repair_attachment: "",
        parts_desc: "",
        series: "",
        operation_type: "",
        area: "",
        find_addr: "",
        find_process: "",
        defect_problem: "",
        station: "",
        a3_principal: "",
        a3_team_top: "",
        a3_team_member: "",
        a3_issue_desc: "",
        a3_contingency: "",
        a3_cause: "",
        a3_affirm: "",
        a3_correct: "",
        a3_verify: "",
        a3_result: "",
        a3_end: "",
        a3_setting: "",
        a3_step: 0,
        a3_issues: '',
        a3_img_base64: ''
        //处置方法

    }
    configure.value = ([
        {
            title: "1.发现问题",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 1
        },
        {
            title: "2.问题描述",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 2
        },
        {
            title: "3.应急响应措施",
            principal: "",
            limit_date: 0,
            department: "",
            issue_number: 3
        },
        {
            title: "4.原因分析",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 4
        },
        {
            title: "5.根本原因确认",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 5
        },
        {
            title: "6.纠正措施",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 6
        },
        {
            title: "7.措施验证",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 7
        },
        {
            title: "8.保持改进结果",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 8
        }, {
            title: "9.传递",
            principal: "",
            department: "",
            limit_date: 0,
            issue_number: 9
        }
    ])
}

const closeDialog = () => {
    initForm()
    drawer.value = false
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (setp.value === 4 && Issues.value.length === 0) {
            ElMessage({
                type: 'error',
                message: '请选择问题在提交',
                showClose: true
            })
            return
        }
        if (Issues.value !== undefined || Issues.value.length > 0) {
            form.value.a3_issues = JSON.stringify(Issues.value)
        }
        form.value.a3_issue_desc = JSON.stringify(tableIssueData.value)
        form.value.a3_cause = JSON.stringify(tableCauseData.value)
        form.value.a3_result = JSON.stringify(tableResultData.value)
        form.value.a3_correct = JSON.stringify(tableCorrectData.value)
        form.value.a3_contingency = JSON.stringify(tableEmergencyData.value)
        form.value.a3_end = JSON.stringify(tableEndData.value)
        form.value.a3_verify = JSON.stringify(tableDeliveryData.value)
        form.value.a3_setting = JSON.stringify(configure.value)
        form.value.a3_step = setp.value
        if (isPass.value == 4) {
            form.value.a3_step = 4
        }
        if (valid) {
            const res = await updateManage(form.value)
            if (res.code === 0) {
                ElMessage({
                    type: 'success',
                    message: '报告提交成功',
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

.setting_body {
    padding: 20px;

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
            display: flex;
            align-items: center;
            justify-content: center;
            flex-direction: column;
            margin-right: 20px;

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