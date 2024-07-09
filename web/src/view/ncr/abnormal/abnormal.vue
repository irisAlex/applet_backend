<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="部门" style="width:200px" prop="method">
                    <el-select v-model="searchInfo.department" placeholder="选择部门">
                        <el-option v-for="item in departmentList" :key="item.authorityId" :label="item.authorityName"
                            :value="item.authorityName">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="类型" style="width:200px" prop="method">
                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                        <el-option v-for="item in moldList" :key="item.value" :label="item.label" :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="类别" prop="method" style="width:200px">
                    <el-select v-model="searchInfo.category" placeholder="请选择">
                        <el-option v-for="item in genreList1" :key="item.name" :label="item.genre" :value="item.genre">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="受检物名称" prop="method" style="width:200px">
                    <el-input v-model="searchInfo.checkout_name" placeholder="受检物名称" />
                </el-form-item>
                <el-form-item label="受检物号" prop="method" style="width:200px">
                    <el-input v-model="searchInfo.checkout_name" placeholder="受检物号" />
                </el-form-item>
                <el-form-item label="处理方式" prop="method" style="width:200px">
                    <el-select v-model="searchInfo.process_mode" placeholder="请选择">
                        <el-option v-for="item in methodOptions" :key="item.value" :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="项目名称" prop="project" style="width:200px">
                    <el-input v-model="searchInfo.project" placeholder="项目名称" />
                </el-form-item>
                <el-form-item label="填表日期" prop="checkout_date" style="width:200px">
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
            </div>
            <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
                <el-table-column align="left" label="ID" min-width="150" prop="ID" sortable="custom" />
                <el-table-column align="left" label="编号" min-width="150" prop="serialnumber" sortable="custom" />
                <el-table-column align="left" label="状态" min-width="150" prop="sta" sortable="custom">
                    <template #default="scope">
                        <el-tag :type="info">
                            {{ scope.row.operation_type !== '-1' ? scope.row.operation_type : '异常' }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="部门" min-width="150" prop="department" sortable="custom" />
                <el-table-column align="left" label="类型" min-width="150" prop="mold" sortable="custom" />
                <el-table-column align="left" label="类别" min-width="150" prop="category" sortable="custom" />
                <el-table-column align="left" label="项目" min-width="150" prop="project" sortable="custom" />
                <el-table-column align="left" label="受检物名称" min-width="150" prop="checkout_name" sortable="custom" />
                <el-table-column align="left" label="受检物号" min-width="150" prop="checkout_number" sortable="custom" />
                <el-table-column align="left" label="处理方式" min-width="150" prop="process_mode" sortable="custom" />
                <el-table-column align="left" label="责任部门" min-width="150" prop="duty_department" sortable="custom" />
                <el-table-column align="left" label="检验日期" min-width="150" prop="checkout_date" sortable="custom">
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
                        <el-button icon="view" type="primary" link @click="editApiFunc(scope.row, 'check')"> 查看
                        </el-button>
                        <el-button icon="edit" type="primary" link @click="editApiFunc(scope.row, 'edit')"
                            :disabled="scope.row.operation_type === '-1'">
                            修改</el-button>
                        <!-- <el-tooltip class="item" effect="dark" content="当前操作不被允许" placement="top-end"
                            :disabled="!(scope.row.operation_type != '' && scope.row.operation_type !== '1')">
                            <el-button icon="tools" type="primary" link @click="editApiFunc(scope.row, 'rework')"
                                :disabled="scope.row.operation_type != '' && scope.row.operation_type !== '1'">返工</el-button>
                        </el-tooltip> -->
                        <!-- <el-tooltip class="item" effect="dark" content="当前操作不被允许" placement="top-end"
                            :disabled="!(scope.row.operation_type != '' && scope.row.operation_type !== '2')">
                            <el-button icon="tools" type="primary" link @click="editApiFunc(scope.row, 'repair')"
                                :disabled="scope.row.operation_type != '' && scope.row.operation_type !== '2'">返修</el-button>
                        </el-tooltip> -->
                        <el-button icon="position" type="primary" link @click="setPassDateFunc(scope.row, '让步接收')"
                            :disabled="scope.row.operation_type != '' && scope.row.operation_type !== '让步接收'">让步接收</el-button>
                        <!-- <el-tooltip class="item" effect="dark" content="当前操作不被允许" placement="top-end"
                            :disabled="!(scope.row.operation_type != '' && scope.row.operation_type !== '3')">
                            <el-button icon="ticket" type="primary" link @click="editApiFunc(scope.row, 'parts')"
                                :disabled="scope.row.operation_type != '' && scope.row.operation_type !== '3'">配做</el-button>
                        </el-tooltip> -->
                        <!-- <el-tooltip class="item" effect="dark" content="当前操作不被允许" placement="top-end"
                            :disabled="!(scope.row.operation_type != '' && scope.row.operation_type !== '5')">
                            <el-button icon="remove" type="primary" link @click="editApiFunc(scope.row, 'die')"
                                :disabled="scope.row.operation_type != '' && scope.row.operation_type !== '5'">报废</el-button>
                        </el-tooltip> -->

                        <el-button icon="monitor" type="primary" link @click="editApiFunc(scope.row, 'createNcr')"
                            :disabled="scope.row.is_ncr || scope.row.operation_type === '让步接收' || scope.row.operation_type === '创建ncr'">创建ncr</el-button>
                        <el-button icon="edit" type="primary" link
                            :disabled="scope.row.is_ncr || (scope.row.operation_type === '让步接收' || scope.row.operation_type === '创建ncr')"
                            @click="editApiFunc(scope.row, 'report')">填写方案</el-button>
                        <el-button icon="delete" type="primary" link @click="deleteApiFunc(scope.row)"
                            :disabled="scope.row.operation_type === '-1'">删除</el-button>
                        <el-button icon="edit" type="primary" link @click="editApiFunc(scope.row, 'review')"
                            :disabled="scope.row.operation_type === '创建ncr'">填写评审</el-button>
                        <el-button icon="circle-close" type="primary" link
                            @click="closeAll(scope.row)">NCR关闭</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>

        </div>

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="80%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <el-form-item label="编号:" prop="serialnumber" style="width:20%">
                    <p v-if="isNcr">{{ form.serialnumber }}</p>
                    <el-input placeholder="编号" size="mini" v-model="form.serialnumber" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="发现部门:" prop="department" style="width:20%">
                    <span v-show="isNcr">{{ form.department }}</span>

                    <el-select v-model="form.department" placeholder="北京安新" style="width:100%" v-if="isNcrDisabled">
                        <el-option v-for="item in departmentList" :key="item.authorityName"
                            :label="`${item.authorityName}`" :value="item.authorityName" />
                    </el-select>
                </el-form-item>
                <el-form-item label="类别:" prop="category" style="width:20%">
                    <span v-show="isNcr">{{ form.category }}</span>

                    <el-select v-model="form.category" placeholder="请选择" style="width:100%" v-if="isNcrDisabled">
                        <el-option v-for="item in genreList1" :key="item.genre" :label="`${item.genre}`"
                            :value="item.genre" />
                    </el-select>
                </el-form-item>
                <el-form-item label="项目:" prop="project" style="width:20%">
                    <span v-show="isNcr">{{ form.project }}</span>

                    <el-select v-model="form.project" placeholder="请选择" style="width:100%" v-if="isNcrDisabled">
                        <el-option v-for="item in projectList" :key="item.name" :label="`${item.name}`"
                            :value="item.name" />
                    </el-select>
                </el-form-item>
                <el-form-item label="产品:" prop="pid" style="width:20%">
                    <span v-show="isNcr">{{ form.pid }}</span>
                    <el-select v-model="form.pid" placeholder="请选择" style="width:100%" v-if="isNcrDisabled">
                        <el-option v-for="item in productList" :key="item.name" :label="`${item.name}`"
                            :value="item.ID" />
                    </el-select>
                </el-form-item>
                <el-form-item label="受检物名称:" prop="checkout_name" style="width:20%">
                    <span v-show="isNcr">{{ form.checkout_name }}</span>

                    <el-input placeholder="受检物名称" size="mini" v-model="form.checkout_name" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="受检物号:" prop="checkout_number" style="width:20%">
                    <span v-show="isNcr">{{ form.checkout_number }}</span>

                    <el-input placeholder="受检物号" size="mini" v-model="form.checkout_number" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="图纸号:" prop="graph_number" style="width:20%">
                    <span v-show="isNcr">{{ form.graph_number }}</span>

                    <el-input placeholder="图纸号" size="mini" v-model="form.graph_number" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="版本号:" prop="version_number" style="width:20%">
                    <span v-show="isNcr">{{ form.version_number }}</span>

                    <el-input placeholder="版本号" size="mini" v-model="form.version_number" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="采购订单号:" prop="purchase_order" style="width:20%">
                    <span v-show="isNcr">{{ form.purchase_order }}</span>

                    <el-input placeholder="采购订单号" size="mini" v-model="form.purchase_order" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="生产订单号:" prop="production_order" style="width:20%">
                    <span v-show="isNcr">{{ form.production_order }}</span>

                    <el-input placeholder="生产订单号" size="mini" v-model="form.production_order" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="发货单号:" prop="delivery_order" style="width:20%">
                    <span v-show="isNcr">{{ form.delivery_order }}</span>

                    <el-input placeholder="发货单号" size="mini" v-model="form.delivery_order" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="收货数量:" prop="packages_number" style="width:20%">
                    <span v-show="isNcr">{{ form.packages_number }}</span>

                    <el-input placeholder="收货数量" size="mini" v-model.number="form.packages_number"
                        v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="货物拒收数量:" prop="reject_packages_number" style="width:20%">
                    <span v-show="isNcr">{{ form.reject_packages_number }}</span>

                    <el-input placeholder="货物拒收数量" size="mini" v-model.number="form.reject_packages_number"
                        v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="样品检验数量:" prop="sample_checkout_number" style="width:20%">
                    <span v-show="isNcr">{{ form.sample_checkout_number }}</span>

                    <el-input placeholder="样品检验数量" size="mini" v-model.number="form.sample_checkout_number"
                        v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="样品拒收数量:" prop="reject_sample_checkout_number" style="width:20%">
                    <span v-show="isNcr">{{ form.reject_sample_checkout_number }}</span>

                    <el-input placeholder="样品拒收数量" size="mini" v-model.number="form.reject_sample_checkout_number"
                        v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="责任部门:" prop="duty_department" style="width:20%">
                    <span v-show="isNcr">{{ form.duty_department }}</span>

                    <el-select v-model="form.duty_department" placeholder="北京安新" style="width:100%"
                        v-if="isNcrDisabled">
                        <el-option v-for="item in departmentList" :key="item.authorityName"
                            :label="`${item.authorityName}`" :value="item.authorityName" />
                    </el-select>


                    <!-- <span v-show="isNcr">{{ form.duty_department }}</span> -->

                    <!-- <el-input placeholder="责任部门" size="mini" v-model="form.duty_department" v-if="isNcrDisabled" /> -->
                </el-form-item>
                <el-form-item label="供应商:" prop="supplier" style="width:20%">
                    <span v-show="isNcr">{{ form.supplier }}</span>

                    <el-select v-model="form.supplier" placeholder="供应商" style="width:100%" v-if="isNcrDisabled">
                        <el-option v-for="item in supplierList" :key="item.name" :label="`${item.name}`"
                            :value="item.name" />
                    </el-select>
                </el-form-item>
                <!-- 异常发现 -->
                <el-form-item label="区域/位置:" prop="area" style="width:20%">
                    <span v-show="isNcr">{{ form.area }}</span>

                    <el-input placeholder="区域/位置" size="mini" v-model="form.area" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="工序/工位:" prop="station" style="width:20%">
                    <span v-show="isNcr">{{ form.station }}</span>

                    <el-input placeholder="工序/工位" size="mini" v-model="form.station" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="发现地点:" prop="find_addr" style="width:20%">
                    <span v-show="isNcr">{{ form.find_addr }}</span>

                    <el-input placeholder="发现地点" size="mini" v-model="form.find_addr" v-if="isNcrDisabled" />
                </el-form-item>
                <el-form-item label="发现工序:" prop="find_process" style="width:20%">
                    <span v-show="isNcr">{{ form.find_process }}</span>

                    <el-input placeholder="发现工序" size="mini" v-model="form.find_process" v-if="isNcrDisabled" />
                </el-form-item>
                <!-- 
                <el-form-item label="标签:" prop="photograph" style="width:100%">
                    <el-image v-show="isNcr" v-for="item in imgList"
                        style="width: 100px; height: 100px;display: block;margin: 5px;box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);"
                        :src="path + item.url" :preview-src-list="[path + item.url]">
                    </el-image>
                    <el-upload action="/api/fileUploadAndDownload/upload" multiple :limit="2" :file-list="fileList"
                        :on-success="handleSuccess" show-file-list="false" :on-remove="handleRemove"
                        v-if="isNcrDisabled">
                        <el-button size="small" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                    </el-upload>
                </el-form-item> -->
                <el-form-item label="缺陷描述:" prop="defect_problem" style="width:20%">
                    <span v-show="isNcr">{{ form.defect_problem }}</span>

                    <el-input placeholder="缺陷描述" size="mini" v-model="form.defect_problem" v-if="isNcrDisabled" />
                </el-form-item>
                <!-- <el-form-item label="问题描述:" prop="describe" style="width:80%">
                    <span v-show="isNcr">{{ form.describe }}</span>

                    <el-input type="textarea" placeholder="请输入内容" v-model="form.describe" maxlength="50" show-word-limit
                        :rows="10" v-if="isNcrDisabled" />
                </el-form-item> -->
                <!-- 异常发现 -->
                <el-form-item label="检验日期:" prop="checkout_date" style="width:23%">
                    <span v-show="isNcr">{{ formatDate(form.checkout_date) }}</span>

                    <el-date-picker v-model="form.checkout_date" type="date" placeholder="选择日期"
                        value-format="YYYY-MM-DDT15:04:05Z" v-if="isNcrDisabled">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="问题描述:" prop="describe" style="width:80%;">
                    <!-- <span v-show="isNcr">{{ form.describe }}</span> -->
                    <!-- <el-input type="textarea" placeholder="请输入内容" v-model="form.describe" maxlength="50" show-word-limit
                        :rows="10" v-if="isNcrDisabled" /> -->
                    <div v-if="isNcr">
                        <QuillEditor :options="editorOptions" content-type="html" ref="quillEditor" theme="snow"
                            v-model:content="form.describe" :value="form.describe" readOnly />
                    </div>
                    <div v-if="isNcrDisabled">
                        <QuillEditor :options="editorOptions" content-type="html" ref="quillEditor" theme="snow"
                            v-model:content="form.describe" :value="form.describe" />
                    </div>
                </el-form-item>
                <el-form-item label="隔离标识:" prop="photograph" style="width:100%">
                    <el-image v-show="isNcr" v-for="item in imgList"
                        style="width: 100px; height: 100px;display: block;margin: 5px;box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);"
                        :src="path + item.url" :preview-src-list="[path + item.url]">
                    </el-image>
                    <el-upload action="/api/fileUploadAndDownload/upload" multiple :limit="2" :file-list="fileList"
                        :on-success="handleSuccess" show-file-list="false" with-credentials="true"
                        :on-remove="handleRemove" v-if="isNcrDisabled" :headers="{ 'x-token': userStore.token }">
                        <el-button size="small" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                    </el-upload>
                </el-form-item>
                <!-- 是否通知其他人 -->
                <el-form-item label="是否通知其他人:" prop="process_mode" style="width:100%" v-if="isCreateNcr">
                    <el-switch v-model="notifiy" active-color="#13ce66" inactive-color="#E0E0E0">
                    </el-switch>
                </el-form-item>
                <!-- 通知 -->
                <div v-if="isEnableNotify">
                    <el-form-item label="通知部门:" prop="notification_department" style="width:20%">
                        <el-select v-model="notification_department" placeholder="选择部门" style="width:100%"
                            @change="departmentUinfo(notification_department)">
                            <el-option v-for="item in departmentList" :key="item.authorityName"
                                :label="`${item.authorityName}`" :value="item.authorityId" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="通知人员:" prop="notification_member" style="width:20%">
                        <el-select v-model="notification_member" multiple placeholder="选择人员" style="width:100%">
                            <el-option v-for="item in AuthorityUsers" :key="item.ID" :label="`${item.userName}`"
                                :value="item.userName" />
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 创建ncr -->
                <div v-if="isCreateNcr">
                    <el-form-item label="类型:" prop="mold" style="width:30%">
                        <!-- <span v-show="isNcr">{{ form.mold }}</span> -->

                        <el-select v-model="form.mold" placeholder="请选择" style="width:100%">
                            <el-option v-for="item in moldList" :key="item.value" :label="`${item.label}`"
                                :value="item.value" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="处理方式:" prop="process_mode" style="width:30%">
                        <!-- <span v-show="isNcr">{{ form.process_mode }}</span> -->

                        <el-select v-model="form.process_mode" placeholder="请选择" style="width:100%">
                            <el-option v-for="item in methodOptions" :key="item.value" :label="`${item.label}`"
                                :value="item.value" />
                        </el-select>
                    </el-form-item>
                    <!-- 具体措施 -->
                    <el-form-item label="采取措施措施:" prop="process_mode" style="width:100%">
                        <el-radio-group v-model="form.operation_type" size="medium">
                            <el-radio-button label="返工"> </el-radio-button>
                            <el-radio-button label="返修"></el-radio-button>
                            <el-radio-button label="配做"></el-radio-button>
                            <el-radio-button label="让步接收"></el-radio-button>
                            <el-radio-button label="报废"></el-radio-button>
                            <el-radio-button label="返回供应商"></el-radio-button>
                            <el-radio-button label="补发"></el-radio-button>
                            <el-radio-button label="索赔"></el-radio-button>
                            <el-radio-button label="变更文件"></el-radio-button>
                        </el-radio-group>
                    </el-form-item>
                </div>
                <!-- rework 返工 -->
                <div v-if="isRework">
                    <el-form-item label="数量" prop="rework_number" style="width:20%">
                        <el-input placeholder="数量" size="mini" v-model.number="form.rework_number" />
                    </el-form-item>
                    <el-form-item label="工时" prop="rework_man_hour" style="width: 20%">
                        <el-input placeholder="工时" size="mini" v-model.number="form.rework_man_hour" />
                    </el-form-item>
                    <el-form-item label="工料" prop="rework_quantities" style="width:20%">
                        <el-input placeholder="工料" size="mini" v-model="form.rework_quantities" />
                    </el-form-item>
                    <el-form-item label="工序" prop="rework_process" style="width:20%">
                        <el-input placeholder="工序" size="mini" v-model="form.rework_process" />
                    </el-form-item>
                    <el-form-item label="返工计划完成时间" prop="rework_plan_date" style="width:40%">
                        <el-date-picker v-model="form.rework_plan_date" type="date" placeholder="选择日期"
                            value-format="YYYY-MM-DDT15:04:05Z">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="返工描述" prop="rework_desc" style="width:100%">
                        <el-input type="textarea" placeholder="请输入内容" v-model="form.rework_desc" maxlength="50"
                            show-word-limit :rows="10" />
                    </el-form-item>
                    <el-form-item label="返工证据" prop="rework_attachment" style="width:100%">
                        <el-upload action="/api/fileUploadAndDownload/upload" multiple :limit="2" :file-list="fileList1"
                            :on-success="handleSuccess1" show-file-list="false" :on-remove="handleRemove1"
                            :headers="{ 'x-token': userStore.token }">
                            <el-button size="small" type="primary">点击上传</el-button>
                            <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                        </el-upload>
                    </el-form-item>
                </div>

                <!-- repair 返修 -->
                <div v-if="isRepair">
                    <el-form-item label="返修计划完成时间" prop="repair_plan_date" style="width:50%">
                        <el-date-picker v-model="form.repair_plan_date" type="date" placeholder="选择日期"
                            value-format="YYYY-MM-DDT15:04:05Z">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="返修描述" prop="repair_desc" style="width:100%">
                        <el-input type="textarea" placeholder="请输入内容" v-model="form.repair_desc" maxlength="50"
                            show-word-limit :rows="10" />
                    </el-form-item>
                    <el-form-item label="返修附件" prop="repair_attachment" style="width:100%">
                        <el-upload action="/api/fileUploadAndDownload/upload" multiple :limit="2" :file-list="fileList2"
                            :on-success="handleSuccess2" show-file-list="false" :on-remove="handleRemove2"
                            :headers="{ 'x-token': userStore.token }">
                            <el-button size="small" type="primary">点击上传</el-button>
                            <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                        </el-upload>
                    </el-form-item>
                </div>

                <!-- 配做  -->
                <div v-if="isParts">
                    <el-form-item label="配做方案" prop="parts_desc" style="width:100%">
                        <div>
                            <QuillEditor :options="editorOptions" content-type="html" ref="quillEditor" theme="snow"
                                v-model:content="form.parts_desc" :value="form.parts_desc" />
                        </div>
                        <!-- <el-input type="textarea" placeholder="请输入内容" v-model="form.parts_desc" maxlength="50"
                            show-word-limit :rows="10" /> -->
                    </el-form-item>
                    <el-form-item prop="series" label="配做订单">
                        <el-table :data="tablePartsData" border style="width: 100%;" :stripe="true">
                            <el-table-column label="产品系列号" align="center" min-width="150" prop="product_serialnumber">
                                <template #default="scope">
                                    <el-input v-model="scope.row.product_serialnumber"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="产品名称" align="center" min-width="150" prop="product_name">
                                <template #default="scope">
                                    <el-input v-model="scope.row.product_name"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="物料系列号" min-width="150" prop="w_serialnumber">
                                <template #default="scope">
                                    <el-input v-model="scope.row.w_serialnumber"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="物料名称" align="center" min-width="150" prop="w_name">
                                <template #default="scope">
                                    <el-input v-model="scope.row.w_name"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="操作" min-width="50" prop="action" align="center">
                                <template #default="scope">
                                    <el-button @click="deleteTableData(scope.row)" link icon="Delete"
                                        type="primary"></el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                        <el-icon @click="addTableData" class="icon" size="24" color="#fb7a14">
                            <CirclePlusFilled />
                        </el-icon>
                    </el-form-item>
                </div>

                <div v-if="isReview">
                    <el-form-item label="保密等级:" prop="privary" style="width:30%">
                        <!-- <span v-show="isNcr">{{ form.mold }}</span> -->

                        <el-select v-model="form.privary" placeholder="请选择" style="width:100%">
                            <el-option v-for="item in privaryList" :key="item.value" :label="`${item.label}`"
                                :value="item.value" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="异常来源:" prop="souce" style="width:30%">
                        <!-- <span v-show="isNcr">{{ form.mold }}</span> -->

                        <el-select v-model="form.souce" placeholder="请选择" style="width:100%">
                            <el-option v-for="item in souceList" :key="item.value" :label="`${item.label}`"
                                :value="item.value" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="技术评审" prop="technology_view" style="width:100%"
                        v-if="tc_pri === userStore.userInfo.userName">
                        <el-input type="textarea" placeholder="请输入内容" v-model="form.technology_view" maxlength="50"
                            show-word-limit :rows="10" />
                    </el-form-item>
                    <el-form-item label="工艺评审" prop="craft_view" style="width:100%"
                        v-if="cr_pri === userStore.userInfo.userName">
                        <el-input type="textarea" placeholder="请输入内容" v-model="form.craft_view" maxlength="50"
                            show-word-limit :rows="10" />
                    </el-form-item>
                </div>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog" v-if="isFormDisabled">取 消</el-button>
                    <el-button type="primary" @click="enterDialog" v-if="isFormDisabled">{{ buttonName
                        }}</el-button>
                    <el-button type="primary" @click="setNcrFunc" v-if="isCNcr">创建NCR</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>
<!-- @click="uploadImg" -->
<script setup>
import {
    getAuthorityList,
    getGenreList,
    getSupplierList,
    getProjectList,
    deleteManage,
    updateManage,
    getManageById,
    createManage,
    getManageList,
    setPassDate,
    setNcr,
    getUserAuthorityList,
    sendMessage,
    getUserByNcrID,
    closeAllByID,
    getProductList,
    getProductById
} from '@/api/manage.js'
import { ref, reactive, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toSQLLine, formatDate } from '@/utils/stringFun'
import { useUserStore } from '@/pinia/modules/user'
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';

defineOptions({
    name: 'Api',
})
const buttonName = ref('提交')
const quillEditor = ref();
const userStore = useUserStore()
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
const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_CLI_PORT + '/'
const isFormDisabled = ref(true)
const isNcr = ref(false)
const isNcrDisabled = ref(true)
const isRework = ref(false)
const isRepair = ref(false)
const isParts = ref(false)
const isCreateNcr = ref(false)
const isCNcr = ref(true)
const isReview = ref(false)
const apis = ref([])
const departmentList = ref([])
const genreList1 = ref([])
const supplierList = ref([])
const projectList = ref([])
const fileList = ref([])
const fileList1 = ref([])
const fileList2 = ref([])
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
    status: "",
    is_ncr: false,
    authority_id: 0,
    departments: '',
    privary: '',
    souce: '',
    technology_view: '',
    craft_view: '',
    pid: 0
})
const notifiy = ref(false)
const isEnableNotify = ref(false)
watch(
    () => notifiy.value,
    () => {
        if (notifiy.value) {
            isEnableNotify.value = true
        } else {
            isEnableNotify.value = false
        }
    }
)

watch(
    () => form.value.operation_type,
    () => {
        if (type.value === 'edit' || type.value == 'review') {
            return
        }
        isFormDisabled.value = true
        isNcr.value = true
        isRework.value = false
        isRepair.value = false
        isParts.value = false
        switch (form.value.operation_type) {
            case '返工':
                dialogTitle.value = '创建/修改返工订单'
                isRework.value = true
                isNcrDisabled.value = false
                type.value = 'rework'
                break
            case '返修':
                dialogTitle.value = '创建/修改返修订单'
                isRepair.value = true
                isNcrDisabled.value = false
                type.value = 'repair'
                break
            case '配做':
                dialogTitle.value = '创建/修改配做订单'
                isParts.value = true
                isNcrDisabled.value = false
                type.value = 'parts'
                break
            default:
                break;
        }
    }
)


watch(
    () => form.value.rework_plan_date,
    () => {
        if (form.value.rework_plan_date === '0001-01-01T00:00:00Z') {
            form.value.rework_plan_date = ''
        }
    }
)



const tablePartsData = ref([])
const addTableData = () => {
    const newRow = {
        product_serialnumber: null,
        product_name: null,
        w_serialnumber: null,
        w_name: null
    }
    tablePartsData.value.push(newRow)
}
// 删除
const deleteTableData = (row) => {
    const index = tablePartsData.value.indexOf(row);
    if (index !== -1) {
        tablePartsData.value.splice(index, 1);
    }
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

const privaryList = ref([
    {
        value: 'Open/开放',
        label: 'Open/开放',
        type: 'success'
    },
    {
        value: 'Internal/内部适用',
        label: 'Internal/内部适用',
        type: ''
    },
    {
        value: 'Confidential/保密',
        label: 'Confidential/保密',
        type: 'warning'
    },
    {
        value: 'StrictlyConfidential/严格保密',
        label: 'StrictlyConfidential/严格保密',
        type: 'warning'
    }
])

const souceList = ref([
    {
        value: 'Supplier',
        label: 'Supplier',
        type: 'success'
    },
    {
        value: 'Production',
        label: 'Production',
        type: ''
    },
    {
        value: 'Design',
        label: 'Design',
        type: 'warning'
    },
    {
        value: 'TransPort',
        label: 'TransPort',
        type: 'warning'
    },
    {
        value: 'Packing',
        label: 'Packing',
        type: 'warning'
    },
    {
        value: 'Stock',
        label: 'Stock',
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

//返工图片处理
const handleSuccess1 = (resp) => {
    if (resp.code === 0) {
        ElMessage({
            type: 'success',
            message: '图片上传成功',
            showClose: true
        })

        fileList1.value.push({ name: resp.data.file.name, url: resp.data.file.url })
        form.value.rework_attachment = JSON.stringify(fileList1.value)
        return
    }
    ElMessage({
        type: 'error',
        message: '图片上传失败',
        showClose: true
    })

};

const handleRemove1 = (file, fileList1) => {
    // 处理删除文件的逻辑，例如从文件列表中删除文件
    const index = fileList1.indexOf(file);
    if (index !== -1) {
        fileList.splice(index, 1);
    }
    form.value.rework_attachment = JSON.stringify(fileList1.value)
};
//返修图片处理
const handleSuccess2 = (resp) => {
    if (resp.code === 0) {
        ElMessage({
            type: 'success',
            message: '图片上传成功',
            showClose: true
        })

        fileList2.value.push({ name: resp.data.file.name, url: resp.data.file.url })
        form.value.repair_attachment = JSON.stringify(fileList2.value)
        return
    }
    ElMessage({
        type: 'error',
        message: '图片上传失败',
        showClose: true
    })
};

const handleRemove2 = (file, fileList2) => {
    // 处理删除文件的逻辑，例如从文件列表中删除文件
    const index = fileList2.indexOf(file);
    if (index !== -1) {
        fileList2.splice(index, 1);
    }
    form.value.repair_attachment = JSON.stringify(fileList2.value)
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

const AuthorityUsers = ref([])
const notification_department = ref('')
const notification_member = ref('')
//获取部门用户
const departmentUinfo = async (Authority) => {
    const table = await getUserAuthorityList({ page: page.value, pageSize: pageSize.value, authority_id: Authority, ...searchInfo.value })
    if (table.code === 0) {
        console.log(table.data.list)

        AuthorityUsers.value = table.data.list
    }
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

const productList = ref([])
//产品
const product = async () => {
    const table = await getProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        productList.value = table.data.list
    }
}

product()

const productInfo = ref()
const tc_pri = ref('')
const cr_pri = ref('')
//getProductInfo
const getProductInfo = async () => {
    if (form.value.pid <= 0) {
        return
    }
    const table = await getProductById({ id: form.value.pid })
    if (table.code === 0) {
        productInfo.value = table.data.product
    }
    let p = JSON.parse(productInfo.value.configuration)
    tc_pri.value = p[0].principal
    cr_pri.value = p[1].principal
}


// 查询
const getTableData = async () => {
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

const deleteVisible = ref(false)

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
        status: "",
        is_ncr: false,
        privary: '',
        souce: '',
        technology_view: '',
        craft_view: '',
        pid: 0
        //处置方法

    }
}

const dialogTitle = ref('创建NCR')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
    isFormDisabled.value = true
    isNcr.value = true
    isRework.value = false
    isRepair.value = false
    isParts.value = false
    isReview.value = false
    switch (key) {
        case 'addApi':
            dialogTitle.value = '异常发现'
            isNcr.value = false
            isCNcr.value = false
            break
        case 'edit':
            dialogTitle.value = '编辑异常'
            buttonName.value = '提交'
            isNcr.value = false
            isNcrDisabled.value = true
            isCNcr.value = false
            isFormDisabled.value = true
            break
        case 'report':
            dialogTitle.value = '填写工单'
            isNcr.value = true
            isCNcr.value = false
            isCreateNcr.value = true
            isFormDisabled.value = true
            isNcrDisabled.value = false
            if (form.value.operation_type == '返工') {
                isRework.value = true
            }
            if (form.value.operation_type == '返修') {
                isRepair.value = true
            }
            if (form.value.operation_type == '配做') {
                isParts.value = true
            }
            break
        case 'review':
            dialogTitle.value = '填写评审'
            isNcr.value = true
            isCNcr.value = false
            isCreateNcr.value = false
            isFormDisabled.value = true
            isNcrDisabled.value = false
            isReview.value = true
            break
        case 'check':
            dialogTitle.value = '查看NCR'
            isFormDisabled.value = false
            isNcr.value = true
            isCNcr.value = false
            isNcrDisabled.value = false
            break
        case 'createNcr':
            buttonName.value = '保存草稿'
            dialogTitle.value = '创建NCR'
            isNcrDisabled.value = false
            isFormDisabled.value = true
            isCreateNcr.value = true
            isNcr.value = true
            isCNcr.value = true
            if (form.value.operation_type == '返工') {
                isRework.value = true
            }
            if (form.value.operation_type == '返修') {
                isRepair.value = true
            }
            if (form.value.operation_type == '配做') {
                isParts.value = true
            }
            break
        default:
            break
    }
    type.value = key
    dialogFormVisible.value = true
}


const closeDialog = () => {
    initForm()
    fileList.value = []
    fileList1.value = []
    fileList2.value = []
    tablePartsData.value = []
    isNcrDisabled.value = true
    isNcr.value = false
    isRework.value = false
    isRepair.value = false
    isParts.value = false
    isCreateNcr.value = false
    dialogFormVisible.value = false
}

const imgList = ref()

const editApiFunc = async (row, operation) => {
    const res = await getManageById({ id: row.ID })
    form.value = res.data.manage
    if (form.value.photograph !== '') {
        fileList.value = JSON.parse(form.value.photograph)
        imgList.value = fileList.value
    }

    if (form.value.rework_attachment !== '') {
        fileList1.value = JSON.parse(form.value.rework_attachment)
    }
    if (form.value.repair_attachment !== '') {
        fileList2.value = JSON.parse(form.value.repair_attachment)
    }

    if (form.value.series !== '') {
        tablePartsData.value = JSON.parse(form.value.series)
    }

    if (operation === 'report') {
        const res = await getUserByNcrID({ ncr_id: row.ID, reportname: userStore.userInfo.userName })
        if (res.code !== 0) {
            return
        }
    }

    if (form.value.rework_plan_date === '0001-01-01T00:00:00Z') {
        form.value.rework_plan_date = ''
    }

    if (form.value.repair_plan_date === '0001-01-01T00:00:00Z') {
        form.value.repair_plan_date = ''
    }

    if (operation === 'review') {
        getProductInfo()
    }

    openDialog(operation)

}

const closeAll = async (row) => {
    const res = await closeAllByID({ id: row.ID })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: 'ncr 流程已关闭',
            showClose: true
        })
    }
    getTableData()
}



const setPassDateFunc = async (row, ot) => {
    ElMessageBox.confirm('是否要进行让步接收操作？, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    })
        .then(async () => {
            const res = await setPassDate({ id: row.ID, operation_type: ot })
            if (res.code === 0) {
                ElMessage({
                    type: 'success',
                    message: '放行更新成功'
                })
                getTableData()
                closeDialog()
            }
        })
}

const setNcrFunc = async (row) => {
    const res = await setNcr({ id: form.value.ID })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: 'NCR创建成功',
            showClose: true
        })
    }
    enterDialog()
    getTableData()
    closeDialog()
}

const sentNotice = async (id) => {
    console.log(notification_member.value)
    for (let item of notification_member.value) {
        console.log(item)
        const res = await sendMessage({ message_type: '1', message_link: '', message_is_active: true, message_receive_name: item, message_content: '有新通知请处理', ncr_id: id })

        console.log(res)
    }
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (tablePartsData.value !== null) {
            form.value.series = JSON.stringify(tablePartsData.value)
        }


        if (form.value.rework_plan_date === '') {
            form.value.rework_plan_date = '0001-01-01T00:00:00Z'
        }

        if (form.value.repair_plan_date === '') {
            form.value.repair_plan_date = '0001-01-01T00:00:00Z'
        }


        if (isEnableNotify.value) {
            sentNotice(form.value.ID)
        }

        if (valid) {
            switch (type.value) {
                case 'addApi':
                    {
                        const res = await createManage(form.value)
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
                case 'rework':
                case 'repair':
                case 'createNcr':
                case 'parts':
                case 'review':
                    {
                        const res = await updateManage(form.value)
                        if (res.code === 0) {
                            ElMessage({
                                type: 'success',
                                message: '草稿提交成功',
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
    ElMessageBox.confirm('此操作将永久删除该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    })
        .then(async () => {
            const res = await deleteManage(row)
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
<style lang='scss' scoped>
.box {
    position: relative;

    .icon {
        position: absolute;
        bottom: 10px;
        right: 19px;
    }
}

:deep(.ql-editor) {
    min-height: 180px;
}

:deep(.ql-formats) {
    height: 21px;
    line-height: 21px;
}
</style>