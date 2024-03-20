import { defineStore } from 'pinia'
import { MailSectionType } from '@/enums/MailSection'
import type { MailListType } from '@/types/MailList'
import { formatDate } from '@/utils/FormatDateTime'
import type { EmailContentType } from '@/types/MailContent'
import { fetchData } from '@/api'

export const useMailingStore = defineStore({
  id: 'mail',
  state: () => ({
    mailSectionType: MailSectionType.INBOX as MailSectionType,
    mailList: [] as MailListType[],
    totalMails: 0 as number,
    currentPage: 1 as number,
    displayPage: 1 as number,
    lastPage: 0 as number,
    currentMail: null as EmailContentType | null,
    searchWord: undefined as string | undefined,
  }),
  actions: {
    setMailListData(data: MailListResponseType) {
      this.totalMails = data.hits.total.value
      this.mailList = data.hits.hits.map(mailResponse => {
        return {
          mailId: mailResponse._id,
          from: mailResponse._source.from,
          subject: mailResponse._source.subject,
          date: formatDate(mailResponse._source.date),
          viewed: false,
        }
      })
      this.lastPage = this.totalMails / 100
    },
    setCurrentMail(index: number, data: MailListResponseType) {
      if (data.hits.total.value == 1) {
        this.mailList[index].viewed = true
        const { from, to, subject, content, date } = data.hits.hits[0]._source
        this.currentMail = {
          from,
          to: to!,
          subject,
          content: content!,
          date: formatDate(date)
        }
      }
    },
    async previousPage() {
      this.currentPage--
      const data = await fetchMailsListByPage(this.mailSectionType, this.currentPage - 1)
      if (data) {
        this.setMailListData(data)
      }
    },
    async nextPage() {
      this.currentPage++
      const data = await fetchMailsListByPage(this.mailSectionType, this.currentPage - 1)
      if (data) {
        this.setMailListData(data)
      }
    },
    async searchMailBy(term: string) {
      const response: MailListResponseType = await fetchData(`/${this.mailSectionType.toLowerCase()}?search=${term}`)
      if (response) {
        this.setMailListData(response)
        this.searchWord = term
      }
    }
  }
})

const fetchMailsListByPage = async (sectionType: MailSectionType, page: number): Promise<MailListResponseType | undefined> => {
  try {
    const endpointUrl = `/${sectionType.toLowerCase()}?page=${page}`
    return await fetchData(endpointUrl)
  } catch (error) {
    console.error('Error fetching data:', error)
  }
}

interface MailResponseType {
  _id: string
  _score: number
  '@timestamp': string
  _source: {
    date: string
    from: string
    message_id: string
    subject: string,
    to?: string,
    content?: string
  }
}

interface HitsType {
  total: {
    value: number
  },
  hits: MailResponseType[]
}

interface MailListResponseType {
  hits: HitsType
}