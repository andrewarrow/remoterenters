require "securerandom"

class Building < ApplicationRecord
  before_create :set_guid

  def set_guid
    self.guid = SecureRandom.uuid
  end
end
