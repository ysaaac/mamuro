import { defineStore } from 'pinia'
import { MailSectionType } from '@/enums/MailSection'
import type { MailListType } from '@/types/MailList'
import { formatDate } from '@/utils/FormatDateTime'
import type { EmailContentType } from '@/types/MailContent'

export const useMailingStore = defineStore({
  id: 'mail',
  state: () => ({
    mailSectionType: MailSectionType.INBOX as MailSectionType,
    mailList: [] as MailListType[],
    totalMails: 0 as number,
    currentMail: null as EmailContentType | null
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
          viewed: false
        }
      })
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
        console.log({ index, currentMail: this.currentMail })
      }
    }
  }
})

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