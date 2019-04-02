<template>
    <div class="record">
        <Banner @pageSwitch="pageSwitch"></Banner>
        <PageAddRecord
            @pageSwitch="pageSwitch"
            @getRecord="getRecord"
            :id="id"
            v-if="pageAddRecordState"
        ></PageAddRecord>
        <div class="recordList">
            <div class="oneRecord" v-for="(item, i) in recordList" :key="i">
                <span class="oneRecord-detail">{{item.detail}}</span>
                <span class="oneRecord-comment">{{item.comment}}</span>
                <button class="btnDel" @click="deleteRecord(item._id)">Del</button>
            </div>
        </div>
    </div>
</template>

<script>
import Banner from "../common/Banner";
import PageAddRecord from "./PageAddRecord";
import request from "../request";
export default {
    data: function() {
        return {
            id: "",
            pageAddRecordState: false,
            recordList: []
        };
    },
    components: { Banner, PageAddRecord },
    methods: {
        pageSwitch: function(b) {
            this.pageAddRecordState = b;
        },
        deleteRecord(id) {
            request.delRecord(id).then(response => {
                this.getRecord();
            });
        },
        getRecord() {
            request.getRecord(this.id).then(response => {
                const data = response.data;
                this.recordList = data || [];
            });
        }
    },
    mounted: function() {
        this.id = this.$route.query.id;
        this.getRecord();
    }
};
</script>

<style scoped>
.recordList {
    max-width: 500px;
    margin: auto;
    display: flex;
    flex-direction: column;
}

.oneRecord {
    background-color: white;
    padding: 10px;
    border-radius: 5px;
    margin: 5px;
    position: relative;
}

.oneRecord-detail {
    display: block;
}

.oneRecord-comment {
    display: block;
    color: gray;
}

.btnDel {
    outline: none;
    background-color: lightcoral;
    color: white;
    cursor: pointer;
    position: absolute;
    bottom: 10%;
    right: 5%;
    padding: 5px 10px;
}
</style>

