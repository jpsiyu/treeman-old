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
            <Record v-for="item in recordList" :key="item._id" :item="item" @getRecord="getRecord"></Record>
        </div>
    </div>
</template>

<script>
import Banner from "../common/Banner";
import PageAddRecord from "./PageAddRecord";
import Record from "./Record";
import request from "../request";
export default {
    data: function() {
        return {
            id: "",
            pageAddRecordState: false,
            recordList: []
        };
    },
    components: { Banner, PageAddRecord, Record },
    methods: {
        pageSwitch: function(b) {
            this.pageAddRecordState = b;
        },
        getRecord() {
            request.getRecord(this.id).then(response => {
                const data = response.data || [];
                data.sort((a, b) => {
                    return b.timestamp - a.timestamp;
                });
                this.recordList = data
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
</style>

