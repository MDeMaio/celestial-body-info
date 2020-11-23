<template>
<div>
    <search-bar @search="searchName" @reset="resetPage"> </search-bar>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <h4 style="font-size: 35px;">Available Stars</h4>
            <div class="btn-group dropleft mb-2">
                <button type="button" class="btn btn-secondary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    Star Classification
                </button>
                <div class="dropdown-menu">
                    <button class="dropdown-item" @click="classification = 'All'" :class="{'active':(classification === 'All')}">All</button>
                    <button class="dropdown-item" @click="classification = 'Yellow Dwarf'" :class="{'active':(classification === 'Yellow Dwarf')}">Yellow Dwarf</button>
                    <button class="dropdown-item" @click="classification = 'Yellow Supergiant'" :class="{'active':(classification === 'Yellow Supergiant')}">Yellow Supergiant</button>
                </div>
            </div>
            <ul class="list-group">
                <li class="list-group-item lgi-pointer py-1" :class="{ active: index == currentIndex }" v-for="(star, index) in stars" :key="star.star_id" @click="setActiveStar(star, index)">
                    {{ star.name }}
                </li>
            </ul>
            <nav aria-label="Pagination">
                <ul class="pagination mt-2 justify-content-center">
                    <li v-if="loaded" :key="'prev'" :class="{'disabled':currentPage === 1 || totalPages === 0}" class="page-item previous-item">
                        <button class="page-link" @click="currentPage--">Prev</button>
                    </li>
                    <li v-for="page in pageArray" :key="page" class="page-item" :class="{'active':(currentPage === page)}">
                        <button class="page-link" @click="currentPage = page">{{page}}</button>
                    </li>
                    <li v-if="loaded" :key="'next'" :class="{'disabled':currentPage === totalPages || totalPages === 0}" class="page-item next-item">
                        <button class="page-link" @click="currentPage++">Next</button>
                    </li>
                </ul>
            </nav>
        </div>
    </div>
    <transition name="slide-fade">
        <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentStar">{{ currentStar.name }}</h1>
    </transition>
    <div class="row">
        <transition name="slide-fade">
            <div class="col-md-4" v-if="currentStar">
                <h1>Basic Information</h1>
                <ul class="list-group">
                    <li :key="'alternate_name'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Mass:</strong> {{currentStar.basic_information.mass}} kg</p>
                    </li>
                    <li :key="'classification'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Classification:</strong> {{currentStar.basic_information.classification}}</p>
                    </li>
                    <li :key="'number_of_satelites'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Radius:</strong> {{currentStar.basic_information.radius}} mi</p>
                    </li>
                    <li :key="'star_system'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>System:</strong> {{currentStar.basic_information.system}}</p>
                    </li>
                    <li :key="'most_abundant_resource'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Temperature:</strong> {{currentStar.basic_information.temperature}} K</p>
                    </li>
                    <li :key="'surface_gravity'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Age:</strong> {{currentStar.basic_information.age}} Years</p>
                    </li>
                </ul>
            </div>
        </transition>
        <transition name="slide-fade">
            <img class="img-center col-md-4 justify-content-right" v-if="currentStar" v-bind:src="'data:image/jpeg;base64,' + currentStar.image">
        </transition>
        <div class="col-md-4">
        </div>
    </div>
    <div class="row mt-3 justify-content-center">
        <div class="col-md-10">
            <div v-if="currentStar">
                <ul class="list-group list-group-flush">
                    <li class="list-group-item mt-2" v-for="(val, index) in currentStar.facts" :key="index">
                        <h1 style="font-size: 40px; text-align: center;">{{val.title}}</h1>
                        <p style="font-size: 22px;">{{val.fact}}</p>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import StarService from "../services/StarService";

export default {
    name: "stars-list",
    data() {
        return {
            stars: [],
            currentStar: null,
            currentIndex: -1,
            pageArray: [],
            currentPage: typeof this.$route.query.page == "undefined" ? 1 : parseInt(this.$route.query.page),
            totalPages: 0,
            classification: typeof this.$route.query.page == "undefined" ? "All" : this.$route.query.classification,
            loaded: false,
            resetting: false,
            searching: false
        };
    },
    methods: {
        retrieveStars(page, classification, name) { // Fetchs all of our stars for the current page.
            StarService.getAll(page, classification, name == "" ? "All" : name)
                .then(response => {
                    this.stars = response.data.stars;
                    this.currentPage = page;
                    this.classification = classification;
                    this.pageArray = this.generatePaginationPageArray(response.data.number_of_documents, 5);
                    this.clearStarView();
                })
                .catch(e => {
                    console.log(e);
                }).finally(() => {
                    this.resetting = false; // why here????
                    this.searching = false; // ???????????
                });
        },

        clearStarView() { // Refreshes the page to the default state.
            this.currentIndex = -1;
            this.currentStar = null;
        },

        setActiveStar(star, index) { // Updates currently viewed star.
            this.currentStar = star;
            this.currentIndex = index;
        },

        searchName(value) { // This will search for a star by name and return it, caps specific as of now.
            this.searching = true;
            this.retrieveStars(1, "All", value);
            this.$router.push({
                query: Object.assign({}, this.$route.query, {
                    page: 1,
                    classification: "All",
                    name: value
                })
            }).catch(() => {}); // Catch the error because we dont care about redirecting to same page.;

        },
        validatePageCount(count, total) {
            if (count > total) {
                count = total;
            }
            if (count % 2 == 0) { //ensure odd number
                count--;
            }
            return count;
        },
        generatePaginationPageArray(numOfDocuments, recordsPerPage) { // Fetch our pages in an array so we can iterate over it later to create the pagination list items.
            if (numOfDocuments == 0) { // No paging if no records.
                this.totalPages = 0;
                return;
            }

            const pageArray = [];
            this.totalPages = Math.ceil(numOfDocuments / recordsPerPage); // sets total page count for data
            let pageDisplayCount = this.validatePageCount(5, this.totalPages);
            const mid = Math.ceil(pageDisplayCount / 2); // sets mid of displayed pages
            let startIndex = (this.currentPage + 1) - mid; // set first page to be displayed

            if (this.totalPages <= recordsPerPage) {
                startIndex = 1;
                pageDisplayCount = this.totalPages;
            } else if (this.currentPage < mid) { // account for front half
                startIndex = 1;
            } else if ((this.currentPage + (mid - 1)) > this.totalPages) { //account for back half
                startIndex = this.totalPages - pageDisplayCount + 1;
            }
            for (let i = startIndex; i < startIndex + pageDisplayCount; i++) {
                pageArray.push(i);
            }
            return pageArray;
        },

        resetPage() {
            this.resetting = true;
            this.retrieveStars(1, 'All', 'All');
            this.$router.push(this.$route.path).catch(() => {}); // Catch the error because we dont care about redirecting to same page.
        }
    },
    watch: { // Watch for data change in which page the user is currently on, call API to get new data when it changes.
        "currentPage": function () {
            if (this.resetting || this.searching) {
                return;
            }

            const name = typeof this.$route.query.name == "undefined" ? "All" : this.$route.query.name;
            this.$router.push({
                query: Object.assign({}, this.$route.query, {
                    page: this.currentPage,
                    classification: this.classification,
                    name: name
                })
            });
            this.retrieveStars(this.currentPage, this.classification, name);
        },

        "classification": function () {
            if (this.resetting || this.searching) {
                return;
            }

            this.$router.push({
                query: Object.assign({}, this.$route.query, {
                    page: this.currentPage,
                    classification: this.classification,
                    name: "All"
                })
            });
            if (this.currentPage === 1) { // Otherwise changing current page will take care of the refresh for us.
                this.retrieveStars(this.currentPage, this.classification, "All");
                return;
            }
            this.currentPage = 1;
        },

        "stars": function (val) {
            if (val === null) {
                return;
            }

            if (val.length === 1) {
                this.currentStar = val[0];
                this.currentIndex = 0;
            }
        }
    },
    mounted() {
        const name = typeof this.$route.query.name == "undefined" || this.$route.query.name == "All" ? "" : this.$route.query.name
        // Todo: Load drop-down menu items here for classifications.
        this.retrieveStars(this.currentPage, this.classification, name);
        this.loaded = true;
    }
};
</script>

<style>
.lgi-pointer {
    cursor: pointer;
}

.lgi-pointer:hover {
    background-color: #53a5fc;
}

.lgi-colored-space {
    background-color: #fcfbfe;
    color: #1d1135;
}

.img-center {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
}

.fs-20 {
    font-size: 20px;
}

/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
    transition: all 1s ease;
}

.slide-fade-leave-active {
    transition: all 1s ease;
}

.slide-fade-enter,
.slide-fade-leave-to

/* .slide-fade-leave-active below version 2.1.8 */
    {
    transform: translatey(300px);
    opacity: 0;
}
</style>
