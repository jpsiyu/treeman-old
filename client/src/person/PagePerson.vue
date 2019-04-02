<template>
    <div class="pagePerson">
        <Banner @pageSwitch="pageSwitch"></Banner>
        <div class="pagePerson-list">
            <Member v-for="info in members" :key="info.id" :info="info" @getAllPerson="getAllPerson"/>
        </div>
        <PageAddPerson v-if="pageAddState" @pageSwitch="pageSwitch" @getAllPerson="getAllPerson"></PageAddPerson>
    </div>
</template>

<script>
import { mapState } from "vuex";
import request from "../request";
import Member from "./Member";
import Banner from "../common/Banner";
import PageAddPerson from "./PageAddPerson";
export default {
    name: "PagePerson",
    data: function() {
        return {
            pageAddState: false
        };
    },
    computed: {
        ...mapState({
            members: state => state.person.members
        })
    },
    methods: {
        pageSwitch: function(b) {
            this.pageAddState = b;
        },
        getAllPerson: function() {
            request.getAllPerson().then(response => {
                const data = response.data || [];
                this.$store.commit("initMembers", data);
            });
        }
    },
    components: { Member, Banner, PageAddPerson },
    mounted: function() {
        this.getAllPerson();
    }
};
</script>

<style scoped>
.pagePerson {
    position: relative;
}

.pagePerson-list {
    display: flex;
    margin: auto;
    justify-content: center;
    flex-direction: column;
    padding: 10px;
}

@media screen and (min-width: 768px) {
    .pagePerson-list {
        padding: 10px 0;
        flex-direction: row;
        justify-content: space-between;
        flex-wrap: wrap;
        max-width: 768px;
    }
}
</style>
