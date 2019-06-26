Vue.component('folder-contents', {
  props: ['folder', 'childrenOf'],
  template: `
    <details>
      <summary>📁 {{folder.name}}</summary>
      <ul>
        <li v-for="item in childrenOf(folder)">
          <folder-contents
            v-if="item.folder"
            :folder="item"
            :childrenOf="childrenOf"
          ></folder-contents>
          <file-item :item="item" v-else>
          </file-item>
        </li>
      </ul>
    </details>
  `
})

Vue.component('file-item', {
  props: ['item'],
  template: `
    <span class="file">
      📄 {{item.name}}
    </span>
  `
})

new Vue({
  el: "#vue",
  data: {
    fileItems: [],
    nextFolderId: 0
  },
  computed: {
    homeFolder() {
      return this.fileItems[0]
    },
    homeFolders() {
      return this.fileItems.filter(
        item => item.folder && item.parent === null
      )
    }
  },
  methods: {
    createFolder(name, parent) {
      this.fileItems.push({
        id: this.nextFolderId++,
        folder: true,
        name,
        parent
      })
    },
    createFile(name, parent) {
      this.fileItems.push({
        folder: false,
        name,
        parent
      })
    },
    childrenOf(folder) {
      return this.fileItems.filter(item => item.parent === folder.name)
    }
  },
  created() {
    this.createFolder("myBucket", null)
    this.createFolder("other-bucket", null)
    this.createFolder("examplePublicBlucket", null)
    this.createFolder("publicVideoBucket", null)

    this.createFile("cat.jpg", "examplePublicBlucket")
    this.createFile("facncy-cat.jpg", "examplePublicBlucket")
    this.createFile("crypto-cat.jpg", "examplePublicBlucket")

    this.createFile("fun-video.mp4", "publicVideoBucket")
    this.createFile("josephUpload1.mp4", "publicVideoBucket")
    this.createFile("hugeVideoFile.mp4", "publicVideoBucket")

    this.createFolder("insideBucketFolder", "myBucket")
    this.createFile("mikejson.json", "insideBucketFolder")

    this.createFolder("Secret", "insideBucketFolder")
    this.createFolder("Super Secret", "Secret")
    this.createFolder("Super Duper Secret", "Super Secret")
    this.createFile("passwords.txt", "Super Duper Secret")

  }
})
